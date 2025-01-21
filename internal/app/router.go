package app

import (
	"gouldcs/gouldservice/dao/player"
	"gouldcs/gouldservice/dao/team"
	"gouldcs/gouldservice/dao/utils"
	"gouldcs/gouldservice/internal/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	modules := []Module{
		modules.NewCityModule(db),
	}

	for _, module := range modules {
		module.RegisterRoutes(router)
	}

	router.Group("/api/v1").
		GET("/teams", utils.WithDB(db, team.GetTeams)).
		GET("/teams/:TeamId", utils.WithDB(db, team.GetTeamById)).
		GET("/players", utils.WithDB(db, player.GetPlayers)).
		GET("/players/:PlayerId", utils.WithDB(db, player.GetPlayerById)).
		GET("/teams/:TeamId/players", utils.WithDB(db, player.GetPlayersOnTeam)).
		GET("/cities/:CityId/players", utils.WithDB(db, player.GetPlayersOnTeam)).
		PUT("/teams", utils.WithDB(db, team.PutTeam))
	return router
}