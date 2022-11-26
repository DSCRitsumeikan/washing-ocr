package interactor

import (
	"app/src/entity/model"

	"github.com/line/line-bot-sdk-go/linebot"
)

type messageReplyInteractor struct {
}

func NewMessageReplyInteractor() *messageReplyInteractor {
	return &messageReplyInteractor{}
}

func (interactor *messageReplyInteractor) ReplyMessages(bot *linebot.Client, events []*linebot.Event) (err error) {
	err = model.ReplyMessages(bot, events)
	return err
}
