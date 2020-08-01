package main

import (
	"os"

	controllers "github.com/aizu-go-kapro/keiGo/backend/controllers/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		kc := new(controllers.KeigoController)
		api.GET("/keigo", kc.ConvertKeigo)
	}
	router.Run(":" + port)
}
