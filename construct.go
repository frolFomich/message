package message

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
)

var (
	ErrorInvalidAbstractDocument = errors.New("*AbstractDocument expected but not found")
)

func New(options ...OptMessage) Message {
	msg := &messageImpl{
		*doc.New(),
	}
	for _, opt := range options {
		opt(msg)
	}
	if msg.Id() == "" || msg.Type() == "" || msg.Timestamp().IsZero() {
		panic(ErrorMessageIsInvalid)
	}
	return msg
}

func FromDocument(d doc.Document) Message {
	ad, ok := d.(*doc.AbstractDocument)
	if !ok {
		panic(ErrorInvalidAbstractDocument)
	}
	return &messageImpl{
		*ad,
	}
}
