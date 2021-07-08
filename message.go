package message

import (
	doc "github.com/frolFomich/abstract-document"
	"time"
)

//Message message interface
type Message interface {
	doc.Document
	//Id unique message id
	Id() string
	//Timestamp message creation timestamp
	Timestamp() time.Time
	//Type type of message
	Type() string
	//DataType data type
	DataType() string
	//Data message payload
	Data() interface{}
}
