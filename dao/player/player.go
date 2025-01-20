package player

import (
	"log"
	"net/http"

	"gouldcs/gouldservice/dao/team"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Player struct {
	PlayerID uint `gorm:"primaryKey"`
	Firstname string `gorm:"column:first_name"`
	Lastname string `gorm:"column:last_name"`
	JerseyNumber string
	Team team.Team `gorm:"foreignKey:TeamID"`
}

func GetPlayers(c *gin.Context, db *gorm.DB) {
	var players []Player
	result := db.
				Preload("Team").
				Find(&players)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, players)
}

func GetPlayerById(c *gin.Context, db *gorm.DB) {
	ProvidedPlayerId := c.Param("PlayerId")
	log.Printf("Player Id: %s", ProvidedPlayerId)
	var player Player
	result := db.
				Preload("Team").
				Where("player_id = ?", ProvidedPlayerId).
				First(&player)

	if (result.Error == nil) {
		c.IndentedJSON(http.StatusOK, player)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "player not found"})
	}
}

func GetPlayersOnTeam(c *gin.Context, db *gorm.DB) {
	ProvidedTeamId := c.Param("TeamId")
	var players []Player
	result := db.Where("team_id = ?", ProvidedTeamId).Find(&players)

	if (result.Error == nil) {
		c.IndentedJSON(http.StatusOK, players)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "team not found or has no players"})
	}
}

func GetPlayersInCity(c *gin.Context, db *gorm.DB) {
	ProvidedCityId := c.Param("CityId")
	var players []Player

	result := db.Joins("JOIN cities ON players.team_id = cities.team_id").
				Where("cities.city_id = ?", ProvidedCityId).
				Find(&players)

	if result.Error == nil {
		c.IndentedJSON(http.StatusOK, players)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "city not found or has no players"})
	}
}