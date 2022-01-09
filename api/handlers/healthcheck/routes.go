package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	router := r.Group("/healthcheck")
	router.GET("", get)
}
