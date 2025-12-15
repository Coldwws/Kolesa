package handler

import (
	"sync"

	"github.com/Coldwws/kolesa/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	mu   sync.RWMutex
	data map[int64]models.Car
	lastID int64
}

func NewHandler() *Handler {
	return &Handler{
		data : make(map[int64]models.Car),
	}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	car := router.Group("/")
	{
	car.GET("/cars", h.GetAllCars)
	car.POST("/cars",h.CreateCar)
	car.GET("/cars/:id",h.GetCarByID)
	car.PATCH("/cars/:id",h.UpdateCar)
	car.DELETE("/cars/:id",h.DeleteCar)
	}

	return router
}

