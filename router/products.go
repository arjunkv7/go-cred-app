package router

import (
	"go-cred-app/handlers"
	"go-cred-app/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductsRouter(router *gin.Engine) {
	router.GET("/products", middlewares.AuthMiddleware(), handlers.GetProductFromDB)
	router.POST("/add-product", middlewares.AuthMiddleware(), handlers.AddProductToDB)
}
