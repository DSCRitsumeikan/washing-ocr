package main

import (
	"app/src/application/interactor"
	"app/src/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	sampleInteractor := interactor.NewSampleInteractor()
	sampleController := controller.NewSampleController(sampleInteractor)

	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.GET("/sample", func(c *gin.Context) { sampleController.ShowSample(c) })

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed r.Run: %w", err)
	}
}
