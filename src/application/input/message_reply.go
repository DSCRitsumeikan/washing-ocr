package input

type MessageReplyUsecase interface {
	ReplyMessages(Events) error
}

type Message interface {
	Message()
}

type Event struct {
	ReplyToken string
	Type       string
	Message    Message
}

type Events []Event
