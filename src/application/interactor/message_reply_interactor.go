package interactor

import (
	"app/src/entity/repository"
)

type messageReplyInteractor struct {
	replyMessageHandler repository.MessageReplyRepositoy
}

func NewMessageReplyInteractor(replyMessageHandler repository.MessageReplyRepositoy) *messageReplyInteractor {
	return &messageReplyInteractor{
		replyMessageHandler: replyMessageHandler,
	}
}

func (interactor *messageReplyInteractor) ReplyMessages() (err error) {
	err = interactor.replyMessageHandler.ReplyMessages()
	return err
}
