package event

type Reply interface {
	Event
	Items() interface{}
	HasMore() interface{}
}

type ReplayReply interface {
	Reply
	SubjectId() string
}
