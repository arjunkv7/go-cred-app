package router

import (
	handlers "go-cred-app/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	
	{
		authRoutes.POST("/login", handlers.LoginHandler)
		authRoutes.POST("/register", handlers.RegisterHandler)
	}

}
