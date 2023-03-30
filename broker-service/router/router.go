package router

import (
	"broker/internal/payload"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(payload *payload.Handler) {
	// Options
	r = gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	h := r.Group("/api/v1")
	{
		//h.POST("/handler", userHandler.CreateUser)
		h.GET("/broker", payload.Broker)
		h.POST("/handler", payload.Authenticate)
	}
}

func Start(addr string) error {
	return r.Run(addr)
}
