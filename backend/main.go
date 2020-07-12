package main

import (
	"io/ioutil"

	controllers "github.com/aizu-go-kapro/keiGo/backend/controllers/v1"
	"github.com/aizu-go-kapro/keiGo/backend/utils"
	"github.com/gin-gonic/gin"
)

// go fmt ...

func main() {
	kenjo, err := ioutil.ReadFile("./utils/" + "kenjo.json")
	if err != nil {
		panic(err)
	}
	util := utils.Utils{
		Kenjo: kenjo,
	}
	router := gin.Default()
	kC := controllers.NewKeigo()
	api := router.Group("/api/v1")
	{
		api.GET("/keigo", kC.ConvertKeigo)
	}
	router.Run(":3000")
}
