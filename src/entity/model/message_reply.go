package model

type Message interface {
	Message()
}

type EventType string

const EventTypeMessage EventType = "message"

type MessageReplyEvent struct {
	ReplyToken string
	Type       EventType
	Message    Message
}

type MessageReplyEvents []MessageReplyEvent
