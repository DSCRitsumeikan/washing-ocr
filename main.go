package main

import (
	"app/src/application/interactor"
	"app/src/controller"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadenv()
	sampleInteractor := interactor.NewSampleInteractor()
	sampleController := controller.NewSampleController(sampleInteractor)
	messageReplyInteractor := interactor.NewMessageReplyInteractor()
	messageReplyController := controller.NewMessageReplyController(messageReplyInteractor)

	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.GET("/sample", func(c *gin.Context) { sampleController.ShowSample(c) })
	v1.POST("/reply", func(c *gin.Context) { messageReplyController.ReplyMessage() })

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed r.Run: %w", err)
	}
}

func loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}
