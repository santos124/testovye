package client_route_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	store2 "starline_test/internal/store"
	"sync/atomic"
	"time"

)
type Location struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Type string  `json:"type"`
}

type CostingOptions struct {
	Auto `json:"auto"`
}

type Auto struct {
	UseTolls          float64 `json:"use_tolls"`
	UseHighways       int     `json:"use_highways"`
	UseTracks         float64 `json:"use_tracks"`
	UseDistance       float64 `json:"use_distance"`
	ServicePenalty    int     `json:"service_penalty"`
	ServiceFactor     int     `json:"service_factor"`
	ManeuverPenalty   int     `json:"maneuver_penalty"`
	Width             int     `json:"width"`
	Height            int     `json:"height"`
}

type DateTime struct {
	Type int `json:"type"`
}

type DirectionsOptions struct {
	Units    string `json:"units"`
	Format   string `json:"format"`
	Language string `json:"language"`
}

type RouteRequest struct {
	Locations        []Location       `json:"locations"`
	Costing          string           `json:"costing"`
	Alternates       int              `json:"alternates"`
	CostingOptions   CostingOptions   `json:"costing_options"`
	DateTime         DateTime         `json:"date_time"`
	DirectionsOptions DirectionsOptions `json:"directions_options"`
}
type RouteResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry string  `json:"geometry"`
		Legs     []struct {
			Distance float64 `json:"distance"`
			Duration float64 `json:"duration"`
			Steps    []struct {
				CongIdx  []float64 `json:"cong_idx"`
				Distance float64   `json:"distance"`
				Duration float64   `json:"duration"`
				Geometry string    `json:"geometry"`
				Lengths  []int     `json:"lengths"`
				Maneuver struct {
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Instruction   string    `json:"instruction"`
					IsToll        bool      `json:"is_toll"`
					Location      []float64 `json:"location"`
					Type          string    `json:"type"`
				} `json:"maneuver"`
				Name          string   `json:"name"`
				SegmentShapes []string `json:"segment_shapes"`
				SpeedLimits   []int    `json:"speed_limits"`
				Destinations  string   `json:"destinations,omitempty"`
			} `json:"steps"`
			Summary      string        `json:"summary"`
			ViaWaypoints []interface{} `json:"via_waypoints"`
			Weight       float64       `json:"weight"`
		} `json:"legs"`
		UseToll    bool    `json:"use_toll"`
		Weight     float64 `json:"weight"`
		WeightName string  `json:"weight_name"`
	} `json:"routes"`
	UUID      string `json:"uuid"`
	Waypoints []struct {
		Distance float64   `json:"distance"`
		Location []float64 `json:"location"`
		Name     string    `json:"name"`
	} `json:"waypoints"`
}

type duration_driver struct {
	duration float64
	driverID int
}

func FindNearestDriver(store *store2.Store, lat, lon float64) int {
	store.RLock()
	defer store.RUnlock()
	cnt := int32(0)
	end := make(chan struct{})
	chanRes := make(chan duration_driver)
	nearestDriverID := -1
	go catchResult(chanRes, &nearestDriverID, end)
	for driverID, driver := range store.Mapa {
		if atomic.LoadInt32(&cnt) > 3 {
			for atomic.LoadInt32(&cnt) > 3 {
				time.Sleep(time.Millisecond*10)
				continue
			}
		}
		go func (latfrom, lonfrom, latto, lonto float64, drID int)  {
			atomic.AddInt32(&cnt, 1)
			defer atomic.AddInt32(&cnt, -1)
			// Вызов функции для получения ETA от точки клиента до точки водителя
			duration, err := getETADuration(latfrom, lonfrom, latto, lonto)
			if err != nil {
				log.Println("Failed to get ETA:", err)
				return
			}
			chanRes <- duration_driver{
				duration: float64(duration),
				driverID: drID,
			}	
		}(lat, lon, driver.Latitude, driver.Longitude, driverID)
	}
	close(chanRes)
	<-end
	return nearestDriverID
}

func catchResult(ch chan duration_driver, nearestDriverID *int, end chan struct{}) {
	mapa := map[int]float64{}
	cnt := 0
	for dd := range ch {
		mapa[dd.driverID] = dd.duration
		cnt++
		if cnt % 10 == 0 {
			log.Println("cnt:", cnt)
		}

	}
	minDuration := float64(0)
	for id, duration := range mapa {
		if minDuration == 0 || duration < minDuration {
			minDuration = duration
			*nearestDriverID = id
		}
	}
	close(end)
}

func getETADuration(fromLat, fromLon, toLat, toLon float64) (float64, error) {
	url := "api_starline"

	
	str := RouteRequest{
		Locations: []Location{
			{Lat: fromLat, Lon: fromLon, Type: "break"},
			{Lat: toLat, Lon: toLon, Type: "break"},
		},
		Costing:    "auto",
		Alternates: 1,
		CostingOptions: CostingOptions{
			Auto: Auto{
				UseTolls:        0.5,
				UseHighways:     1,
				UseTracks:       0.5,
				UseDistance:     0.6,
				ServicePenalty:  75,
				ServiceFactor:   1,
				ManeuverPenalty: 50,
				Width:           3,
				Height:          2,
			},
		},
		DateTime: DateTime{
			Type: 0,
		},
		DirectionsOptions: DirectionsOptions{
			Units:    "km",
			Format:   "osrm",
			Language: "ru-RU",
		},
	}

	dataOut, err := json.Marshal(str)
	if err != nil {
		return 0, err
	}
	client := http.Client{}
	client.Timeout = time.Millisecond * 500
	// log.Fatalf("%s", dataOut)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(dataOut))
	if err != nil {
		return 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()


	data := RouteResponse{}
	dataIN, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(dataIN, &data)
	// log.Fatalf("%s %v %v %v %v", dataIN ,fromLat, fromLon, toLat, toLon)
	if err != nil {
		return 0, err
	}

	if len(data.Routes) == 0 {
		return 0, fmt.Errorf("No routes found")
	}

	return data.Routes[0].Duration, nil
}
