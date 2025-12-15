package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Coldwws/kolesa/internal/models"
	"github.com/Coldwws/kolesa/internal/validate"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCars(c *gin.Context) {
		h.mu.RLock()
		defer h.mu.RUnlock()

		cars := make([]models.Car,0,len(h.data))

		for _,car := range h.data{
			cars = append(cars,car)
		}

		c.JSON(200, cars)
}


func (h *Handler)CreateCar(c *gin.Context){
		var car models.Car

		if err := c.ShouldBindJSON(&car);err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		if ok,msg := validate.ValidateCar(car);!ok{
			c.JSON(http.StatusBadRequest,gin.H{"error":msg})
			return
		}

		h.mu.Lock()
		defer h.mu.Unlock()

		h.lastID++
		car.ID = h.lastID
		h.data[car.ID] = car

		c.JSON(http.StatusCreated,car)

}

func (h *Handler) GetCarByID(c *gin.Context){
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam,10,64)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		h.mu.RLock()
		car, ok := h.data[id]
		h.mu.RUnlock()
		

		if !ok{
			c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,car)
}

func (h *Handler)UpdateCar(c *gin.Context){
	idParam := c.Param("id")
	
	id, err := strconv.ParseInt(idParam,10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	var UpdateCar models.UpdateCar
	if err := c.ShouldBindJSON(&UpdateCar); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if ok,msg := validate.ValidateUpdateCar(UpdateCar); !ok{
		c.JSON(http.StatusBadRequest, gin.H{"error":msg})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	
	car,ok := h.data[id]
	if !ok{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"car not found",
		})
		return
	}
	if UpdateCar.Title != nil{
		car.Title = *UpdateCar.Title
	}
	if UpdateCar.Brand != nil{
		car.Brand = *UpdateCar.Brand
	}
	if UpdateCar.Model != nil{
		car.Model = *UpdateCar.Model
	}
	if UpdateCar.Description != nil{
		car.Description = *UpdateCar.Description
	}
	if UpdateCar.Price != nil{
		car.Price = *UpdateCar.Price
	}
	if UpdateCar.Year != nil{
		car.Year = *UpdateCar.Year
	}
	if UpdateCar.Mileage != nil{
		car.Mileage = *UpdateCar.Mileage
	}
	if UpdateCar.EngineType != nil{
		car.EngineType = *UpdateCar.EngineType
	}
	if UpdateCar.EngineVolume != nil{
		car.EngineVolume = *UpdateCar.EngineVolume
	}
	if UpdateCar.Transmission != nil{
		car.Transmission = *UpdateCar.Transmission
	}
	if UpdateCar.DriveType != nil{
		car.DriveType = *UpdateCar.DriveType
	}
	if UpdateCar.BodyType != nil{
		car.BodyType = *UpdateCar.BodyType
	}
	if UpdateCar.Color != nil{
		car.Color = *UpdateCar.Color
	}
	if UpdateCar.Steering != nil{
		car.Steering = *UpdateCar.Steering
	}
	if UpdateCar.City != nil{
		car.City = *UpdateCar.City
	}
	if UpdateCar.Status != nil{
		car.Status = *UpdateCar.Status
	}
	car.UpdatedAt = time.Now()
	h.data[id] = car

	c.JSON(http.StatusOK,car)
}

func (h *Handler)DeleteCar(c *gin.Context){
	idParam := c.Param("id")
	id,err := strconv.ParseInt(idParam,10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.data[id];!ok{
		c.JSON(http.StatusNotFound, gin.H{"error":"car not found"})
		return
	}
	delete(h.data,id)
	c.Status(http.StatusNoContent)

}