package aerrors

import (
	"errors"
	"reflect"
	"sync"
)

var errorPool = &sync.Pool{
	New: func() interface{} {
		return &AError{
			buf: make([]byte, 0, 500),
		}
	},
}

func newAError(code Code, reason string) *AError {
	e := errorPool.Get().(*AError)
	e.buf = e.buf[:0]
	e.withCode(code).withReason(reason)
	return e
}

func putEvent(e *AError) {
	errorPool.Put(e)
}

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

type Error interface {
	Error() string
}

type Builder interface {
	WithParent(parent error) Builder
	WithMessage(message string) Builder
	WithStack() Builder
	Err() Error
	withCode(code Code) Builder
	withReason(reason string) Builder
}

type AError struct {
	parent  error
	code    Code
	reason  string
	message string
	stack   string
	buf     []byte
}

func New(code Code, reason string) Builder {
	return newAError(code, reason)
}

func (err *AError) WithParent(parent error) Builder {
	if err == nil {
		return nil
	}
	err.parent = parent
	err.buf = err.appendString(err.appendKey(err.buf, "parent"), parent.Error())
	return err
}

func (err *AError) WithMessage(message string) Builder {
	if err == nil {
		return nil
	}
	err.message = message
	err.buf = err.appendString(err.appendKey(err.buf, "message"), message)
	return err
}

func (err *AError) WithStack() Builder {
	if err == nil {
		return nil
	}
	const depth = 32
	err.stack = LogStack(2, depth)
	return err
}

func (err *AError) Err() Error {
	if err == nil {
		return nil
	}
	err.buf = err.appendString(err.appendLineBreak(err.buf), err.stack)
	putEvent(err)
	return err
}

// nolint
func (err AError) Error() string {
	return BytesToString(err.buf)
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

func (err *AError) withCode(code Code) Builder {
	err.code = code
	err.buf = err.appendString(err.appendKey(err.buf, "code"), code.Error())
	return err
}

func (err *AError) withReason(reason string) Builder {
	err.reason = reason
	err.buf = err.appendString(err.appendKey(err.buf, "reason"), reason)
	return err
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

func (err *AError) appendKey(dst []byte, key string) []byte {
	if (len(dst)) != 0 {
		dst = append(dst, ',')
	}
	return append(err.appendString(dst, key), ':')
}

func (err *AError) appendString(dst []byte, s string) []byte {
	b := StringToBytes(s)
	dst = append(dst, b...)
	return dst
}

func (err *AError) appendLineBreak(dst []byte) []byte {
	return append(dst, '\n')
}
