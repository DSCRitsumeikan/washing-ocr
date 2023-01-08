package repository

import "app/src/entity/model"

type MessageReply interface {
	ReplyMessages(events model.MessageReplyEvents) error
}
