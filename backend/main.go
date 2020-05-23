package main

import (
	"github.com/gin-gonic/gin"
	controllers "keigo/controllers/v1"
)

func main() {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		kc := new(controllers.KeigoController)
		api.GET("/keigo", kc.ConvertKeigo)
	}
	router.Run(":3000")
}
