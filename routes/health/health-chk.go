package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func Router(router *gin.Engine) *gin.Engine {
	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return router
}
