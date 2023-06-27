package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	client_route_api "starline_test/internal/client-route-api"
	config2 "starline_test/internal/config"
	store2 "starline_test/internal/store"
	"strconv"
)

type App struct {
	Config  *config2.Config
	Storage *store2.Store
	Gin     *gin.Engine
}

func (app *App) HandleDriverSearch(c *gin.Context) {
	// Извлечение параметров lat и lon из URL-запроса
	latStr := c.Query("lat")
	lonStr := c.Query("lon")

	// Проверка наличия параметров lat и lon
	if latStr == "" || lonStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing latitude or longitude"})
		return
	}

	// Преобразование параметров lat и lon в тип float64
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	// Поиск ближайшего водителя
	driverID := client_route_api.FindNearestDriver(app.Storage, lat, lon)

	c.JSON(http.StatusOK, gin.H{"driver_id": driverID})
}
