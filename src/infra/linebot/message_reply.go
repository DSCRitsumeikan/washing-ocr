package linebot

import (
	"app/src/entity/repository"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type replyMessageHandler struct {
	bot    *linebot.Client
	events []*linebot.Event
}

func NewReplyMessageHandler() (repository.MessageReplyRepositoy, error) {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	var c *gin.Context

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Println(err)
		}
		return nil, fmt.Errorf("line bot invalid signature: %v", err)
	}

	return &replyMessageHandler{
			bot:    bot,
			events: events},
		nil
}

func (handler replyMessageHandler) ReplyMessages() (err error) {
	bot := handler.bot
	events := handler.events

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch event.Message.(type) {
			case *linebot.ImageMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("未完成")).Do(); err != nil {
					log.Println(err)
				}
			}
		}
	}
	return err
}
