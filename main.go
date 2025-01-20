package main

import (
	"gouldcs/gouldservice/dao/city"
	"gouldcs/gouldservice/dao/player"
	"gouldcs/gouldservice/dao/team"
	"gouldcs/gouldservice/dao/utils"

	"github.com/gin-gonic/gin"
)

func main() {
    db, err := GetDB()
    if (err != nil) {
        print("Error")
    }
    router := gin.Default()
    router.GET("/teams", utils.WithDB(db, team.GetTeams))
    router.GET("/cities", utils.WithDB(db, city.GetCities))
    router.GET("/players", utils.WithDB(db, player.GetPlayers))
    router.GET("/players/:PlayerId", utils.WithDB(db, player.GetPlayerById))
    router.GET("/teams/:TeamId/players", utils.WithDB(db, player.GetPlayersOnTeam))
    router.GET("/cities/:CityId/players", utils.WithDB(db, player.GetPlayersOnTeam))



    router.Run("localhost:8080")
}