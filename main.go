package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/teams", getTeams)
    router.GET("/players", getPlayers)
	router.GET("/players/:PlayerId", getPlayerById)

    router.Run("localhost:8080")
}