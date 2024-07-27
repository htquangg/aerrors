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
