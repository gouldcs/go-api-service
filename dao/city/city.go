package city

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type City struct {
	CityID uint `gorm:"primaryKey;autoIncrement" json:"city_id"`
	CityName string `gorm:"column:city_name" json:"city_name"`
	CityAlias string `gorm:"column:city_alias" json:"city_alias"`
	CityState string `gorm:"column:city_state" json:"city_state"`
}

func GetCities(c *gin.Context, db *gorm.DB) {
	var cities []City
	result := db.Find(&cities)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, cities)
}

func GetCityById(c *gin.Context, db *gorm.DB) {
	var city City
	ProvidedCityId := c.Param("CityId")
	result := db.Where("city_id = ?", ProvidedCityId).Find(&city)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, city)
}

func PutCity(c *gin.Context, db *gorm.DB) {
	var newCity City

	if err := c.ShouldBindJSON(&newCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	result := db.Create(&newCity)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCity)
}