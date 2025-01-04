package routes

import (
	"github.com/eavlongs/go_backend_template/controllers"
	"github.com/eavlongs/go_backend_template/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterMainRoutes(router *gin.RouterGroup, c *controllers.MainController, m *middlewares.MainMiddleware) {
	routes := router.Group("/main")
	routes.GET("/", c.Test)
}
