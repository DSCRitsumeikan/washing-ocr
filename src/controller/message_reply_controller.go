package controller

import (
	"app/src/application/input"
	"log"
)

type messageReplyController struct {
	usecase input.MessageReplyUsecase
}

func NewMessageReplyController(usecase input.MessageReplyUsecase) messageReplyController {
	return messageReplyController{
		usecase: usecase,
	}
}

func (controller messageReplyController) ReplyMessage() {
	err := controller.usecase.ReplyMessages()
	if err != nil {
		log.Println(err)
	}
}
