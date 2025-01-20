package city

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type City struct {
	CityID uint `gorm:"primaryKey"`
	CityName string `gorm:"column:city_name"`
	CityAlias string `gorm:"column:city_alias"`
	CityState string `gorm:"column:city_state"`
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