package controller

import (
	"app/src/entity/repository"
	"log"
)

type messageReplyController struct {
	repository repository.MessageReplyRepositoy
}

func NewMessageReplyController(repository repository.MessageReplyRepositoy) messageReplyController {
	return messageReplyController{
		repository: repository,
	}
}

func (controller messageReplyController) ReplyMessage() {
	err := controller.repository.ReplyMessages()
	if err != nil {
		log.Println(err)
	}
}
