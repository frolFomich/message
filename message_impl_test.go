package message

import (
	doc "github.com/frolFomich/abstract-document"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		options []OptMessage
	}

	expectedId := "test-id-1"
	expectedType := "/test/type"
	expectedTimestamp := time.Now()

	tests := []struct {
		name string
		args args
		want *messageImpl
	}{
		{name: "Simple message with required properties", args: args{
			options: []OptMessage{Id(expectedId), Type(expectedType), Timestamp(expectedTimestamp)},
		}, want: &messageImpl{
			doc.Of(map[string]interface{}{
				IdPropertyName: expectedId,
				TypePropertyName: expectedType,
				TimestampPropertyName: float64(expectedTimestamp.UnixNano() / int64(time.Millisecond)),
			}),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got.AsPlainMap(), tt.want.AsPlainMap())
			}
		})
	}
}

func Test_messageImpl_Data(t *testing.T) {
	type fields struct {
		Document doc.Document
	}

	expectedData := map[string]interface{}{
		"A" : "B", "C" : 100.0, "D" : true,
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Data getter", fields: fields{
			doc.Of(map[string]interface{}{
				DataPropertyName: expectedData,
			}),
		}, want: expectedData},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &messageImpl{
				tt.fields.Document,
			}
			if got := m.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageImpl_DataType(t *testing.T) {
	type fields struct {
		Document doc.Document
	}

	expectedDataType := "test/data/type"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "DataType getter", fields: fields{
			doc.Of(map[string]interface{}{
				DataTypePropertyName: expectedDataType}),
		}, want: expectedDataType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &messageImpl{
				tt.fields.Document,
			}
			if got := m.DataType(); got != tt.want {
				t.Errorf("DataType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageImpl_Id(t *testing.T) {
	type fields struct {
		Document doc.Document
	}

	expectedId := "test-id-1"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Id getter", fields: fields{
			doc.Of(map[string]interface{}{
				IdPropertyName: expectedId,
			}),
		}, want: expectedId},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &messageImpl{
				tt.fields.Document,
			}
			if got := m.Id(); got != tt.want {
				t.Errorf("Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageImpl_Timestamp(t *testing.T) {
	type fields struct {
		Document doc.Document
	}

	expectedTimestamp := time.Unix(0, time.Now().UnixNano() / int64(time.Millisecond) * int64(time.Millisecond))

	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{name: "Timestamp getter", fields: fields{
			doc.Of(map[string]interface{}{
				TimestampPropertyName: float64(expectedTimestamp.UnixNano() / int64(time.Millisecond)),
			}),
		}, want: expectedTimestamp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &messageImpl{
				tt.fields.Document,
			}
			if got := m.Timestamp(); !got.Equal(tt.want) {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageImpl_Type(t *testing.T) {
	type fields struct {
		Document doc.Document
	}

	expectedType := "test/Type"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Type getter", fields: fields{
			doc.Of(map[string]interface{}{
				TypePropertyName: expectedType,
			}),
		}, want: expectedType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &messageImpl{
				tt.fields.Document,
			}
			if got := m.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
