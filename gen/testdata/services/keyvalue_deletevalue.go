// Code generated by thriftrw v1.5.0
// @generated

package services

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/gen/testdata/exceptions"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type KeyValue_DeleteValue_Args struct {
	Key *Key `json:"key,omitempty"`
}

func (v *KeyValue_DeleteValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Key != nil {
		w, err = v.Key.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Key_Read(w wire.Value) (Key, error) {
	var x Key
	err := x.FromWire(w)
	return x, err
}

func (v *KeyValue_DeleteValue_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *KeyValue_DeleteValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	return fmt.Sprintf("KeyValue_DeleteValue_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Key_EqualsPtr(lhs, rhs *Key) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *KeyValue_DeleteValue_Args) Equals(rhs *KeyValue_DeleteValue_Args) bool {
	if !_Key_EqualsPtr(v.Key, rhs.Key) {
		return false
	}
	return true
}

func (v *KeyValue_DeleteValue_Args) MethodName() string {
	return "deleteValue"
}

func (v *KeyValue_DeleteValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var KeyValue_DeleteValue_Helper = struct {
	Args           func(key *Key) *KeyValue_DeleteValue_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*KeyValue_DeleteValue_Result, error)
	UnwrapResponse func(*KeyValue_DeleteValue_Result) error
}{}

func init() {
	KeyValue_DeleteValue_Helper.Args = func(key *Key) *KeyValue_DeleteValue_Args {
		return &KeyValue_DeleteValue_Args{Key: key}
	}
	KeyValue_DeleteValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		case *InternalError:
			return true
		default:
			return false
		}
	}
	KeyValue_DeleteValue_Helper.WrapResponse = func(err error) (*KeyValue_DeleteValue_Result, error) {
		if err == nil {
			return &KeyValue_DeleteValue_Result{}, nil
		}
		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_DeleteValue_Result.DoesNotExist")
			}
			return &KeyValue_DeleteValue_Result{DoesNotExist: e}, nil
		case *InternalError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_DeleteValue_Result.InternalError")
			}
			return &KeyValue_DeleteValue_Result{InternalError: e}, nil
		}
		return nil, err
	}
	KeyValue_DeleteValue_Helper.UnwrapResponse = func(result *KeyValue_DeleteValue_Result) (err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.InternalError != nil {
			err = result.InternalError
			return
		}
		return
	}
}

type KeyValue_DeleteValue_Result struct {
	DoesNotExist  *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
	InternalError *InternalError                    `json:"internalError,omitempty"`
}

func (v *KeyValue_DeleteValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalError != nil {
		w, err = v.InternalError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_DeleteValue_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DoesNotExistException_Read(w wire.Value) (*exceptions.DoesNotExistException, error) {
	var v exceptions.DoesNotExistException
	err := v.FromWire(w)
	return &v, err
}

func _InternalError_Read(w wire.Value) (*InternalError, error) {
	var v InternalError
	err := v.FromWire(w)
	return &v, err
}

func (v *KeyValue_DeleteValue_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalError, err = _InternalError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.DoesNotExist != nil {
		count++
	}
	if v.InternalError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("KeyValue_DeleteValue_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *KeyValue_DeleteValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	if v.InternalError != nil {
		fields[i] = fmt.Sprintf("InternalError: %v", v.InternalError)
		i++
	}
	return fmt.Sprintf("KeyValue_DeleteValue_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *KeyValue_DeleteValue_Result) Equals(rhs *KeyValue_DeleteValue_Result) bool {
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}
	if !((v.InternalError == nil && rhs.InternalError == nil) || (v.InternalError != nil && rhs.InternalError != nil && v.InternalError.Equals(rhs.InternalError))) {
		return false
	}
	return true
}

func (v *KeyValue_DeleteValue_Result) MethodName() string {
	return "deleteValue"
}

func (v *KeyValue_DeleteValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
