package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Player struct {
	PlayerID string `json:"playerId"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	JerseyNumber string `json:"jerseyNumber"`
	TeamID string `json:"teamId"`
}

var players = []Player {
	{PlayerID: "1", Firstname: "DK", Lastname: "Metcalf", JerseyNumber: "14", TeamID: "1"},
	{PlayerID: "2", Firstname: "Geno", Lastname: "Smith", JerseyNumber: "7", TeamID: "1"},
	{PlayerID: "3", Firstname: "Zach", Lastname: "Charbonnet", JerseyNumber: "26", TeamID: "1"},
	{PlayerID: "4", Firstname: "Justin", Lastname: "Herbert", JerseyNumber: "10", TeamID: "2"},
	{PlayerID: "5", Firstname: "J.K.", Lastname: "Dobbins", JerseyNumber: "27", TeamID: "2"},
	{PlayerID: "6", Firstname: "Ladd", Lastname: "McConkey", JerseyNumber: "15", TeamID: "2"},
	{PlayerID: "7", Firstname: "Matthew", Lastname: "Stafford", JerseyNumber: "9", TeamID: "3"},
	{PlayerID: "8", Firstname: "Cooper", Lastname: "Kupp", JerseyNumber: "10", TeamID: "3"},
	{PlayerID: "9", Firstname: "Kyren", Lastname: "Williams", JerseyNumber: "23", TeamID: "3"},
	{PlayerID: "10", Firstname: "C.J.", Lastname: "Stroud", JerseyNumber: "7", TeamID: "4"},
	{PlayerID: "11", Firstname: "Stefon", Lastname: "Diggs", JerseyNumber: "14", TeamID: "4"},
	{PlayerID: "12", Firstname: "Tank", Lastname: "Dell", JerseyNumber: "3", TeamID: "4"},
	{PlayerID: "13", Firstname: "Aaron", Lastname: "Rodgers", JerseyNumber: "8", TeamID: "5"},
	{PlayerID: "14", Firstname: "Davante", Lastname: "Adams", JerseyNumber: "15", TeamID: "5"},
	{PlayerID: "15", Firstname: "Mike", Lastname: "Williams", JerseyNumber: "18", TeamID: "5"},
	{PlayerID: "16", Firstname: "Daniel", Lastname: "Jones", JerseyNumber: "8", TeamID: "6"},
	{PlayerID: "17", Firstname: "Malik", Lastname: "Nabers", JerseyNumber: "1", TeamID: "6"},
	{PlayerID: "18", Firstname: "Tommy", Lastname: "DeVito", JerseyNumber: "15", TeamID: "6"},
	{PlayerID: "19", Firstname: "Tyreek", Lastname: "Hill", JerseyNumber: "10", TeamID: "7"},
	{PlayerID: "20", Firstname: "Odell", Lastname: "Beckham Jr.", JerseyNumber: "3", TeamID: "7"},
	{PlayerID: "21", Firstname: "Tua", Lastname: "Tagovailoa", JerseyNumber: "1", TeamID: "7"},
}

func getPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}

func getPlayerById(c *gin.Context) {
    ProvidedPlayerId := c.Param("PlayerId")
	foundPlayer := getPlayerWithId(ProvidedPlayerId)

	if (foundPlayer != nil) {
		c.IndentedJSON(http.StatusOK, foundPlayer)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "player not found"})
	}
}

func getPlayerWithId(ProvidedId string) *Player {
	var player *Player
	player = nil

	for _, a := range players {
        if a.PlayerID == ProvidedId {
            player = &a
        }
    }

	return player
}