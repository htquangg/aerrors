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

type AErrorOption func(*AError)

type AError struct {
	code    ErrorCode
	parent  error
	id      string
	reason  string
	message string
	stack   string
}

func New(code ErrorCode, reason string, opts ...AErrorOption) *AError {
	err := &AError{code: code, reason: reason}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func (err *AError) Code() ErrorCode {
	return err.code
}

func (err *AError) Parent() error {
	return err.parent
}

func (err *AError) ID() string {
	return err.id
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

func WithCode(code ErrorCode) AErrorOption {
	return func(err *AError) {
		err.code = code
	}
}

func WithParent(parent error) AErrorOption {
	return func(err *AError) {
		err.parent = parent
	}
}

func WithID(id string) AErrorOption {
	return func(err *AError) {
		err.id = id
	}
}

func WithReason(reason string) AErrorOption {
	return func(err *AError) {
		err.reason = reason
	}
}

func WithMessage(message string) AErrorOption {
	return func(err *AError) {
		err.message = message
	}
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

func (err *AError) WithID(id string) *AError {
	if err == nil {
		return nil
	}
	err.id = id
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

func (err *AError) Is(target error) bool {
	t, ok := target.(*AError)
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

func (e ErrorCode) Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return AError{code: e, parent: err, message: msg}
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case AError:
		return AError{parent: err, message: fmt.Sprintf("%s: %s", msg, err.Error())}
	case TypeCoder:
		return AError{parent: err, message: msg}
	default:
		return AError{parent: err, code: ErrInternal, message: fmt.Sprintf("%s: %s", msg, err.Error())}
	}
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
