package main

import (
	"go-cred-app/config"
	routers "go-cred-app/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	// create new http server
	r := gin.Default()

	routers.AuthRouter(r)
	routers.ProductsRouter(r)
	config.ConnectDB()
	r.Run(":3000")
}
