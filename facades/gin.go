package facades

import (
	"log"

	"github.com/Firhan384/framework/contracts/route"

	"github.com/Firhan384/gin"
)

func Route(driver string) route.Route {
	instance, err := gin.App.MakeWith(gin.RouteBinding, map[string]any{
		"driver": driver,
	})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*gin.Route)
}
