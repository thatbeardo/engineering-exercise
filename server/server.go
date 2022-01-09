package server

import (
	"engineeringexercise/server/inject"

	"github.com/gin-gonic/gin"
)

func SetupRouteHandlers() *gin.Engine {
	r := gin.Default()
	inject.SetupHealthcheckRoutes(&r.RouterGroup)
	return r
}
