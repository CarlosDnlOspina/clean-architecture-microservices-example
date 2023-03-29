package router

import (
	"chat/internal/user"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	// Options
	r = gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	h := r.Group("/v1/login")
	{
		h.POST("/signup", userHandler.CreateUser)
		h.POST("/login", userHandler.Login)
	}
}

func Start(addr string) error {
	return r.Run(addr)
}
