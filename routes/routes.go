package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())

	api := engine.Group("api/stori")

	api.GET("health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.0.1",
		})
	})

	return engine
}
