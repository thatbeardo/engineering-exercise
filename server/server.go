package server

import (
	"engineeringexercise/server/inject"

	"github.com/gin-gonic/gin"
)

type Engine interface {
}

func SetupRouteHandlers() Engine {
	r := gin.Default()
	inject.SetupHealthcheckRoutes(&r.RouterGroup)
	return r
}
