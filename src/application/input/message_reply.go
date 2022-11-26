package input

import "github.com/line/line-bot-sdk-go/linebot"

type MessageReplyUsecase interface {
	ReplyMessages(bot *linebot.Client, events []*linebot.Event) error
}
