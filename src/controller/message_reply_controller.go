package controller

import (
	"app/src/application/input"
	"log"

	"github.com/gin-gonic/gin"
)

type messageReplyController struct {
	usecase input.MessageReplyUsecase
}

func NewMessageReplyController(usecase input.MessageReplyUsecase) messageReplyController {
	return messageReplyController{
		usecase: usecase,
	}
}

func (controller messageReplyController) ReplyMessage(c *gin.Context) {
	err := controller.usecase.ReplyMessages()
	if err != nil {
		log.Println(err)
	}
}
