package team

import (
	"net/http"

	"gouldcs/gouldservice/dao/city"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Team struct {
    TeamID uint `gorm:"primaryKey"`
    City city.City `gorm:"foreignKey:CityID"`
    TeamName string `gorm:"column:team_name"`
}

func GetTeams(c *gin.Context, db *gorm.DB) {
	var teams []Team
	result := db.Find(&teams)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, teams)
}

func GetTeamById(c *gin.Context, db *gorm.DB) {
	var team Team
	ProvidedTeamId := c.Param("TeamId")

	result := db.
				Preload("City").
				Where("team_id = ?", ProvidedTeamId).
				Find(&team)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, team)
}