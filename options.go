package message

import (
	"errors"
	"time"
)

type OptMessage func (msg *messageImpl)

var (
	ErrorMessageIsInvalid = errors.New("message is invalid")
)

func Id(id string) OptMessage {
	return func(msg *messageImpl) {
		if msg != nil && id != "" {
			msg.Put(IdPropertyName, id)
		}
	}
}

func Type(t string) OptMessage {
	return func(msg *messageImpl) {
		if msg != nil && t != "" {
			msg.Put(TypePropertyName, t)
		}
	}
}

func Timestamp(t time.Time) OptMessage {
	return func(msg *messageImpl) {
		if msg != nil && !t.IsZero() {
			msg.Put(TimestampPropertyName, t.UnixNano() / int64(time.Millisecond))
		}
	}
}

func DataType(dt string) OptMessage {
	return func(msg *messageImpl) {
		if msg != nil && dt != "" {
			msg.Put(DataTypePropertyName, dt)
		}
	}
}

func Data(d interface{}) OptMessage {
	return func(msg *messageImpl) {
		if msg != nil {
			msg.Put(DataPropertyName, d)
		}
	}
}

