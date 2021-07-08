package command

import "github.com/frolFomich/message"

type Command interface {
	message.Message
}

type CUDCommand interface {
	Command
	SubjectId() string
}

type CreateCommand interface {
	CUDCommand
	Subject() interface{}
}

type UpdateCommand interface {
	CUDCommand
	Subject() interface{}
}

type DeleteCommand interface {
	CUDCommand
}
