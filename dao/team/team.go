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

func GetTeamsById(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        ProvidedTeamId := c.Param("PlayerId")
        var teams []Team
        result := db.Find(&teams, ProvidedTeamId)

        if (result.Error == nil) {
            c.IndentedJSON(http.StatusOK, teams)
            return
        } else {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "team not found"})
            return
        }
    }
}