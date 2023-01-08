package controller

import (
	"app/src/application/input"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type messageReplyController struct {
	usecase input.MessageReplyUsecase
	bot     *linebot.Client
}

func NewMessageReplyController(usecase input.MessageReplyUsecase, bot *linebot.Client) messageReplyController {
	return messageReplyController{
		usecase: usecase,
		bot:     bot,
	}
}

func (controller *messageReplyController) ReplyMessage(c *gin.Context) {
	req, err := controller.bot.ParseRequest(c.Request)
	if err != nil {
		err = fmt.Errorf("failed bot.ParseRequest: %w", err)
		log.Println(err)
	}
	events := make(input.Events, 0, len(req))
	for _, e := range req {
		event := input.Event{
			ReplyToken: e.ReplyToken,
			Type:       string(e.Type),
			Message:    e.Message,
		}
		events = append(events, event)
	}

	err = controller.usecase.ReplyMessages(events)
	if err != nil {
		err = fmt.Errorf("failed controller.usecase.ReplyMessages: %w", err)
		log.Println(err)
	}
}
