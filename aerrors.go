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

type ErrorCode string

func (err ErrorCode) TypeCode() string {
	if err == ErrOK {
		return ""
	}
	return string(err)
}

func (err ErrorCode) Error() string {
	if err == ErrOK {
		return ""
	}

	return string(err)
}

type AError struct {
	code    ErrorCode
	parent  error
	reason  string
	message string
	stack   string
}

func New(code ErrorCode, reason string) *AError {
	return &AError{code: code, reason: reason}
}

func (err *AError) Code() ErrorCode {
	return err.code
}

func (err *AError) Parent() error {
	return err.parent
}

func (err *AError) Reason() string {
	return err.reason
}

func (err *AError) Message() string {
	return err.message
}

func (err *AError) Stack() string {
	return err.stack
}

func (err *AError) WithCode(code ErrorCode) *AError {
	if err == nil {
		return nil
	}
	err.code = code
	return err
}

func (err *AError) WithParent(parent error) *AError {
	if err == nil {
		return nil
	}
	err.parent = parent
	return err
}

func (err *AError) WithReason(reason string) *AError {
	if err == nil {
		return nil
	}
	err.reason = reason
	return err
}

func (err *AError) WithMessage(message string) *AError {
	if err == nil {
		return nil
	}
	err.message = message
	return err
}

func (err *AError) WithStack() *AError {
	if err == nil {
		return nil
	}
	err.message = LogStack(2, 0)
	return err
}

// nolint:gocritic
func (err AError) Error() string {
	str := bytes.NewBuffer([]byte{})
	fmt.Fprintf(str, "code: %s, ", err.code.Error())
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

func (err *AError) Is(target error) bool {
	t, ok := target.(*AError)
	if !ok {
		return false
	}
	if t.code != ErrOK && t.code != err.code {
		return false
	}
	if t.message != "" && t.message != err.message {
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
