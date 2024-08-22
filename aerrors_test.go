package aerrors

import (
	"errors"
	"fmt"
	"testing"
)

var (
	errExample  = errors.New("fail")
	fakeReason  = "Test errors, but use a somewhat realistic reason length."
	fakeMessage = "Test errors, but use a somewhat realistic message length."
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
	return Internal(fakeReason).
		WithParent(errExample).
		WithMessage(fakeMessage).
		WithStack().
		Err()
}

func BenchmarkInternalWithStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Internal(fakeReason).
				WithMessage(fakeMessage).
				WithParent(errExample).
				WithStack().
				Err()
		}
	})
}

func BenchmarkInternalWithoutStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Internal(fakeReason).
				WithMessage(fakeMessage).
				WithParent(errExample).
				Err()
		}
	})
}

func BenchmarkInternaEmptylWithStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Internal(fakeReason).
				WithStack().
				Err()
		}
	})
}

func BenchmarkInternalEmptyWithoutStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Internal(fakeReason).
				Err()
		}
	})
}

func BenchmarkNewEmptyWithStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			New(ErrInternal, fakeReason).
				WithStack().
				Err()
		}
	})
}

func BenchmarkNewEmptylWithoutStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			New(ErrInternal, fakeReason).
				Err()
		}
	})
}

func BenchmarkNewWithStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			New(ErrInternal, fakeReason).
				WithMessage(fakeMessage).
				WithParent(errExample).
				WithStack().
				Err()
		}
	})
}

func BenchmarkNewlWithoutStack(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			New(ErrInternal, fakeReason).
				WithMessage(fakeMessage).
				WithParent(errExample).
				Err()
		}
	})
}
