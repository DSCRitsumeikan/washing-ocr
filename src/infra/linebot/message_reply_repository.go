package linebot

import (
	"app/src/entity/model"
	"app/src/entity/repository"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

type MessageReplyRepository struct {
	bot *linebot.Client
}

func NewMessageReplyRepository(bot *linebot.Client) repository.MessageReply {
	return &MessageReplyRepository{
		bot: bot,
	}
}

func NewLinebotClient() (*linebot.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		err = fmt.Errorf("failed linebot.New: %w", err)
		return nil, err
	}

	return bot, err
}

func (mrr *MessageReplyRepository) ReplyMessages(events model.MessageReplyEvents) error {
	for _, event := range events {
		if event.Type == model.EventTypeMessage {
			switch event.Message.(type) {
			case *linebot.ImageMessage:
				if _, err := mrr.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("未完成")).Do(); err != nil {
					err := fmt.Errorf("failed mrr.bot.ReplyMessage: %w", err)
					return err
				}
			}
		}
	}
	return nil
}
