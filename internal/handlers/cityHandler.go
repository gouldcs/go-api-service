package handlers

import (
	"gouldcs/gouldservice/internal/models"
	"gouldcs/gouldservice/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(service *services.CityService) *CityHandler {
	return &CityHandler{service: service}
}

func (h *CityHandler) GetCityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("CityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
	}

	city, err := h.service.GetCityById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
        return
	}

	c.JSON(http.StatusOK, city)
}

func (h *CityHandler) CreateCity(c *gin.Context) {
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CreateCity(&city); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create city"})
		return
	}

	c.JSON(http.StatusCreated, city)
}