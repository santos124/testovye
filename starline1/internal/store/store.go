package store

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Store struct {
	Mapa map[int]*Position
	sync.RWMutex
}

type Position struct {
	Longitude float64
	Latitude  float64
}

func GetStore(filename string) (*Store, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("bad filename or filename isnt")
	}

	store := Store{
		Mapa:    map[int]*Position{},
		RWMutex: sync.RWMutex{},
	}

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if "driver_id,latitude,longitude" == line {
			continue
		}
		splitLine := strings.Split(line, ",")
		if len(splitLine) < 3 {
			continue
		}
		id, err := strconv.Atoi(splitLine[0])
		longitude, err2 := strconv.ParseFloat(splitLine[1], 64)
		latitude, err3 := strconv.ParseFloat(splitLine[2], 64)
		if err != nil || err2 != nil || err3 != nil {
			return nil, fmt.Errorf("bad file for line (%v:) %v", i+1, line)
		}
		store.Lock()
		store.Mapa[id] = &Position{
			Longitude: longitude,
			Latitude:  latitude,
		}
		store.Unlock()
	}
	log.Println("store loaded from: ", filename)
	return &store, nil
}
