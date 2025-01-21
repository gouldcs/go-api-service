package modules

import (
	"gouldcs/gouldservice/internal/handlers"
	repositories "gouldcs/gouldservice/internal/repositories/city"
	"gouldcs/gouldservice/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CityModule struct {
	Handler *handlers.CityHandler
}

func NewCityModule(db *gorm.DB) *CityModule {
	repo := repositories.NewCityRepository(db)
	service := services.NewCityService(repo)
	handler := handlers.NewCityHandler(service)

	return &CityModule{Handler: handler}
}

func (m *CityModule) RegisterRoutes(router *gin.Engine) {
	router.Group("/api/v1").
		GET("/cities/:CityId", m.Handler.GetCityById).
		PUT("/cities", m.Handler.CreateCity)
}