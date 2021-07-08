package message

import (
	"github.com/frolFomich/abstract-document"
	"time"
)

//Message message interface
type Message interface {
	abstract_document.Document
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
