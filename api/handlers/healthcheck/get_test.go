package healthcheck_test

import (
	"engineeringexercise/api/handlers/healthcheck"
	"engineeringexercise/internal"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Routed_StatusOK(t *testing.T) {
	engine := gin.Default()
	healthcheck.SetupRoutes(&engine.RouterGroup)
	response, cleanup := internal.PerformRequest(engine, "GET", "/healthcheck", "")
	defer cleanup()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
