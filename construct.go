package message

import doc "github.com/frolFomich/abstract-document"

func New(options... OptMessage) *messageImpl {
	msg := &messageImpl{
		Document: doc.New(),
	}
	for _,opt := range options {
		opt(msg)
	}
	if msg.Id() == "" || msg.Type() == "" || msg.Timestamp().IsZero() {
		panic(ErrorMessageIsInvalid)
	}
	return msg
}

func FromDocument(d doc.Document) *messageImpl {
	return &messageImpl{
		Document: d,
	}
}
