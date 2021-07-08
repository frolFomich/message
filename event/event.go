package event

import "github.com/frolFomich/message"

type Event interface {
	message.Message
}

type CUDEvent interface {
	Event
	SubjectId() string
	CorrelationId() string
}

type CreatedEvent interface {
	CUDEvent
	Subject() interface{}
}

type UpdatedEvent interface {
	CUDEvent
	Subject() interface{}
}

type DeletedEvent interface {
	CUDEvent
}
