package aerrors

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	e := A()
	fmt.Printf("%s\n\n", e)
	fmt.Printf("%v\n\n", e)
}

func A() error {
	return B()
}

func B() error {
	return C()
}

func C() error {
	return Internal("internal server error").
		WithParent(fmt.Errorf("db connection error")).
		WithStack()
}

func BenchmarkNewWithoutStack(b *testing.B) {
	err := fmt.Errorf("db connection error")
	for i := 0; i < b.N; i++ {
		Internal("internal server error").
			WithMessage("hello world").
			WithParent(err).
			WithID("ab6a0")
	}
}

func BenchmarkNewWithStack(b *testing.B) {
	err := fmt.Errorf("db connection error")
	for i := 0; i < b.N; i++ {
		Internal("internal server error").
			WithMessage("hello world").
			WithParent(err).
			WithID("ab6a0").
			WithStack()
	}
}

func BenchmarkRawWithoutStack(b *testing.B) {
	err := fmt.Errorf("db connection error")
	for i := 0; i < b.N; i++ {
		var aerror AError
		(*Error)(&aerror).
			WithCode(ErrInternal).
			WithReason("internal server error").
			WithMessage("hello world").
			WithParent(err).
			WithID("ab6a0")
	}
}

func BenchmarkRawWithStack(b *testing.B) {
	err := fmt.Errorf("db connection error")
	for i := 0; i < b.N; i++ {
		var aerror AError
		(*Error)(&aerror).
			WithCode(ErrInternal).
			WithReason("internal server error").
			WithMessage("hello world").
			WithParent(err).
			WithID("ab6a0").
			WithStack()
	}
}
