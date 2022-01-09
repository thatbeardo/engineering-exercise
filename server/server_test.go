package server_test

import (
	"engineeringexercise/server"
	"engineeringexercise/server/inject"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouteHandlers_Empty_Instantiated(t *testing.T) {
	router := server.SetupRouteHandlers()
	assert.NotNil(t, router)
}

func TestSetupRouteHandlers_Healthcheck_Routed(t *testing.T) {
	defer inject.Reset()
	healthcheckRoutesCalled := 0

	inject.SetupHealthcheckRoutes = func(r *gin.RouterGroup) {
		assert.NotNil(t, r)
		healthcheckRoutesCalled++
	}

	server.SetupRouteHandlers()
	assert.Equal(t, 1, healthcheckRoutesCalled)
}
