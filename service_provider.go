package gin

import (
	"github.com/Firhan384/framework/contracts/config"
	"github.com/Firhan384/framework/contracts/foundation"
	"github.com/Firhan384/framework/contracts/http"
	"github.com/Firhan384/framework/contracts/log"
	"github.com/Firhan384/framework/contracts/validation"
)

const RouteBinding = "goravel.gin.route"

var (
	App              foundation.Application
	ConfigFacade     config.Config
	LogFacade        log.Log
	ValidationFacade validation.Validation
	ViewFacade       http.View
)

type ServiceProvider struct{}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.BindWith(RouteBinding, func(app foundation.Application, parameters map[string]any) (any, error) {
		return NewRoute(app.MakeConfig(), parameters)
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	LogFacade = app.MakeLog()
	ValidationFacade = app.MakeValidation()
	ViewFacade = app.MakeView()

	app.Publishes("github.com/Firhan384/gin", map[string]string{
		"config/cors.go": app.ConfigPath("cors.go"),
	})
}
