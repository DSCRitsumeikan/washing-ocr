package controller

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func ReplyMessage(c *gin.Context) {
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
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch event.Message.(type) {
			case *linebot.ImageMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("未完成")).Do(); err != nil {
					log.Println(err)
				}
			default:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("画像以外のフォーマットは受け付けておりません。")).Do(); err != nil {
					log.Println(err)
				}
			}
		}
	}

}
