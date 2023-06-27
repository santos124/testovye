package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"starline_test/internal/common"
	config2 "starline_test/internal/config"
	store2 "starline_test/internal/store"
)

type DriverPosition struct {
	DriverID  int     `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ETARequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type ETAResponse struct {
	Duration int `json:"duration"`
}

func main() {
	app := common.App{}
	app.Config = config2.GetVars()
	storage, err := store2.GetStore(app.Config.PathToCSV)
	if err != nil {
		log.Fatal(err)
	}
	app.Storage = storage
	app.Gin = gin.Default()

	// Обработчик метода /driverSearch
	app.Gin.GET("/driverSearch", app.HandleDriverSearch)

	if err := app.Gin.Run(":4444"); err != nil {
		log.Fatal(err)
	}
}
