package controller

import (
	"app/src/application/input"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type messageReplyController struct {
	usecase input.MessageReplyUsecase
}

func NewMessageReplyController(usecase input.MessageReplyUsecase) messageReplyController {
	return messageReplyController{
		usecase: usecase,
	}
}

func (controller messageReplyController) ReplyMessage(c *gin.Context) {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Println(err)
		}
		return
	}
	err = controller.usecase.ReplyMessages(bot, events)
	if err != nil {
		log.Println(err)
	}
}
