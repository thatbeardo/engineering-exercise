package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	c.JSON(http.StatusOK, "Service up and running")
}
