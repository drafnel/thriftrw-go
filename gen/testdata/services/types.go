// Code generated by thriftrw

package services

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type InternalError struct {
	Message *string `json:"message,omitempty"`
}

func (v *InternalError) ToWire() wire.Value {
	var fields [1]wire.Field
	i := 0
	if v.Message != nil {
		fields[i] = wire.Field{ID: 1, Value: wire.NewValueString(*(v.Message))}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func (v *InternalError) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Message = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *InternalError) String() string {
	var fields [1]string
	i := 0
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", *(v.Message))
		i++
	}
	return fmt.Sprintf("InternalError{%v}", strings.Join(fields[:i], ", "))
}

func (v *InternalError) Error() string {
	return v.String()
}

type Key string

func (v Key) ToWire() wire.Value {
	x := (string)(v)
	return wire.NewValueString(x)
}

func (v Key) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

func (v *Key) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (Key)(x)
	return err
}
