package main

import (
	"app/src/application/interactor"
	"app/src/controller"
	"app/src/infra/linebot"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	bot, err := linebot.NewLinebotClient()
	if err != nil {
		log.Fatal("failed linebot.NewLinebotClient: ", err)
	}

	sampleInteractor := interactor.NewSampleInteractor()
	sampleController := controller.NewSampleController(sampleInteractor)

	messageReplyRepository := linebot.NewMessageReplyRepository(bot)
	messageReplyInteractor := interactor.NewMessageReplyInteractor(messageReplyRepository)
	messageReplyController := controller.NewMessageReplyController(messageReplyInteractor, bot)

	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.GET("/sample", func(c *gin.Context) { sampleController.ShowSample(c) })
	v1.POST("/reply", func(c *gin.Context) { messageReplyController.ReplyMessage(c) })

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed r.Run: ", err)
	}
}
