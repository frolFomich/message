package message

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
