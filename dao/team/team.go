package team

import (
	"log"
	"net/http"

	"gouldcs/gouldservice/dao/city"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Team struct {
    TeamID uint `gorm:"primaryKey" json:"team_id"`
    CityID uint `gorm:"column:city_id" json:"city_id"`
    TeamName string `gorm:"column:team_name" json:"team_name"`
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

func PutTeam(c *gin.Context, db *gorm.DB) {
	var newTeam Team

	if err := c.ShouldBindJSON(&newTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Check that the provided CityID belongs to an existing team
	// so that the error message is clear.
	var city city.City
	if err := db.First(&city, newTeam.CityID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city_id: City does not exist"})
        } else {
			log.Print("failed to get city for some other reason")
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
	}

	result := db.Create(&newTeam)
    if result.Error != nil {
		log.Print("Failed to insert the new team")
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, newTeam)
}