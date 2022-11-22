package main

import (
	"app/src/application/interactor"
	"app/src/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Hello World")
	sampleInteractor := interactor.NewSampleInteractor()
	sampleController := controller.NewSampleController(sampleInteractor)

	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.GET("/sample", func(c *gin.Context) { sampleController.ShowSample(c) })

	r.Run(":8080")

	log.Println("Hello World")
}

// リクエスト飛ばしてもなんともならない。なぜ？
