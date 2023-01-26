package interactor

import (
	"app/src/application/input"
	"app/src/entity/model"
	"app/src/entity/repository"
	"fmt"
)

type messageReplyInteractor struct {
	mr repository.MessageReply
}

func NewMessageReplyInteractor(mr repository.MessageReply) input.MessageReplyUsecase {
	return &messageReplyInteractor{
		mr: mr,
	}
}

func (interactor *messageReplyInteractor) ReplyMessages(es input.Events) error {
	events := make(model.MessageReplyEvents, 0, len(es))
	for _, e := range es {
		event := model.MessageReplyEvent{
			ReplyToken: e.ReplyToken,
			Type:       model.EventType(e.Type),
			Message:    e.Message,
		}
		events = append(events, event)
	}
	if err := interactor.mr.ReplyMessages(events); err != nil {
		err = fmt.Errorf("failed interactor.mr.ReplyMessages: %w", err)
		return err
	}
	return nil
}
