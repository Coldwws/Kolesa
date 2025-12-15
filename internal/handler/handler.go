package handler

import (
	"github.com/Coldwws/kolesa/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {

	service service.CarService
}

func NewHandler(service service.CarService) *Handler {
	return &Handler{
		service: service,
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

