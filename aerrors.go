package aerrors

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type TypeCoder interface {
	TypeCode() string
}

type Code string

func (err Code) TypeCode() string {
	if err == ErrOK {
		return ""
	}
	return string(err)
}

func (err Code) Error() string {
	if err == ErrOK {
		return ""
	}

	return string(err)
}

type Error AError

type AError struct {
	code    Code
	parent  error
	id      string
	reason  string
	message string
	stack   string
	skip    int
}

func New(code Code, reason string) *Error {
	var err AError
	(*Error)(&err).WithCode(code).WithReason(reason)
	return (*Error)(&err)
}

func (err *Error) WithParent(parent error) *Error {
	if err == nil {
		return nil
	}
	err.parent = parent
	return err
}

func (err *Error) WithID(id string) *Error {
	if err == nil {
		return nil
	}
	err.id = id
	return err
}

func (err *Error) WithCode(code Code) *Error {
	if err == nil {
		return nil
	}
	err.code = code
	return err
}

func (err *Error) WithReason(reason string) *Error {
	if err == nil {
		return nil
	}
	err.reason = reason
	return err
}

func (err *Error) WithMessage(message string) *Error {
	if err == nil {
		return nil
	}
	err.message = message
	return err
}

func (err *Error) WithStack() *Error {
	if err == nil {
		return nil
	}
	const depth = 32
	err.stack = LogStack(2+err.skip, depth)
	return err
}

// nolint:gocritic
func (err Error) Error() string {
	str := bytes.NewBuffer([]byte{})
	fmt.Fprintf(str, "code: %s, ", err.code.Error())
	str.WriteString("id: ")
	str.WriteString(err.id + ", ")
	str.WriteString("reason: ")
	str.WriteString(err.reason + ", ")
	str.WriteString("message: ")
	str.WriteString(err.message)
	if err.parent != nil {
		str.WriteString(", error: ")
		str.WriteString(err.parent.Error())
	}
	if err.stack != "" {
		str.WriteString("\n")
		str.WriteString(err.stack)
	}

	return str.String()
}

func (err *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	if t.code != ErrOK && t.code != err.code {
		return false
	}
	if t.reason != "" && t.reason != err.reason {
		return false
	}
	if t.parent != nil && !errors.Is(err.parent, t.parent) {
		return false
	}

	return true
}

func (err *AError) As(target interface{}) bool {
	_, ok := target.(**AError)
	if !ok {
		return false
	}
	reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
	return true
}

func TypeCode(err error) string {
	if err == nil {
		return ErrOK.TypeCode()
	}
	var e TypeCoder
	if errors.As(err, &e) {
		return e.TypeCode()
	}
	return ErrUnknown.TypeCode()
}
