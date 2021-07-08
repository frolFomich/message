package command

type Request interface {
	Command
	Parameters() map[string]interface{}
}

type RequestStatusRequest interface {
	Request
	RequestId() string
}

type ReplayRequest interface {
	Request
	SubjectId() string
}
