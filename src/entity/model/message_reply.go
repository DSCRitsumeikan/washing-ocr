package model

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func ReplyMessages(bot *linebot.Client, events []*linebot.Event) (err error) {
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
	return err
}
