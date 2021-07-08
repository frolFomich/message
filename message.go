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

//Subscriber - messages receiver
type Subscriber interface {
	//OnMessage - invoked when message come to the input
	//  topic - topic from which message come
	//  m - message
	OnMessage(topic string, m Message) bool
}

//Publisher - producer of messages
type Publisher interface {
	//Subscribe - adds new subscriber to topic
	//  topic - interested topic
	//  s - subscriber which will be invoked on new message in topic
	Subscribe(topic string, s Subscriber)
}
