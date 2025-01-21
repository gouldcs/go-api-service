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
    router.GET("/teams/:TeamId", utils.WithDB(db, team.GetTeamById))
    router.GET("/cities", utils.WithDB(db, city.GetCities))
    router.GET("/cities/:CityId", utils.WithDB(db, city.GetCityById))
    router.GET("/players", utils.WithDB(db, player.GetPlayers))
    router.GET("/players/:PlayerId", utils.WithDB(db, player.GetPlayerById))
    router.GET("/teams/:TeamId/players", utils.WithDB(db, player.GetPlayersOnTeam))
    router.GET("/cities/:CityId/players", utils.WithDB(db, player.GetPlayersOnTeam))

    router.PUT("/cities", utils.WithDB(db, city.PutCity))
    router.PUT("/teams", utils.WithDB(db, team.PutTeam))

    router.Run("localhost:8080")
}