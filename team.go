package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team struct {
	TeamID string `json:"teamId"`
	CityID string `json:"cityId"`
	TeamName string `json:"teamName"`
}

var teams = []Team {
	{TeamID: "1", CityID: "1", TeamName: "Seahawks"},
	{TeamID: "2", CityID: "2", TeamName: "Chargers"},
	{TeamID: "3", CityID: "2", TeamName: "Rams"},
	{TeamID: "4", CityID: "3", TeamName: "Texans"},
	{TeamID: "5", CityID: "4", TeamName: "Jets"},
	{TeamID: "6", CityID: "4", TeamName: "Giants"},
	{TeamID: "7", CityID: "5", TeamName: "Dolphins"},
}

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}