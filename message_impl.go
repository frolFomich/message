package message

import (
	"github.com/frolFomich/abstract-document"
	"time"
)

const (
	//IdPropertyName - message id property name
	IdPropertyName = "id"
	//TimestampPropertyName - message timestamp property name
	TimestampPropertyName = "timestamp"
	//TypePropertyName - message type property name
	TypePropertyName = "type"
	//DataTypePropertyName - message data type property name
	DataTypePropertyName = "dataType"
	//DataPropertyName - message data property name
	DataPropertyName = "data"
)

type messageImpl struct {
	abstract_document.AbstractDocument
}

func (m *messageImpl) Id() string {
	v, err := m.String(IdPropertyName)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *messageImpl) Timestamp() time.Time {
	v, err := m.Integer(TimestampPropertyName)
	if err != nil {
		panic(err)
	}
	return time.Unix(0, v*int64(time.Millisecond))
}

func (m *messageImpl) Type() string {
	v, err := m.String(TypePropertyName)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *messageImpl) DataType() string {
	v, err := m.String(DataTypePropertyName)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *messageImpl) Data() interface{} {
	return m.Get(DataPropertyName)
}
