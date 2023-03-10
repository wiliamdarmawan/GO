package pkg

import "fmt"

type i1 interface {
	String() int
}

type i2 interface {
	String() string
}

type i3 interface {
	bar() int
}

type i4 interface {
	String() int
	bar() int
}

func fn() {
	var v1 i1
	_ = v1.(i2) //@ diag(`impossible type assertion; i1 and i2 contradict each other`)
	_ = v1.(i3)
	_ = v1.(i4)
	_ = v1.(fmt.Stringer) //@ diag(`impossible type assertion; i1 and fmt.Stringer contradict each other`)
	_ = v1.(interface {   //@ diag(re`i1 and.+String.+contradict each other`)
		String() string
	})
}

type ExampleType[T uint32 | uint64] interface {
	SomeMethod() T
}

func Fn1[T uint32 | uint64]() {
	var iface ExampleType[uint32]
	_ = iface.(ExampleType[T])
}

func Fn2[T uint64]() {
	// In theory we should flag this, but we don't.
	var iface ExampleType[uint32]
	_ = iface.(ExampleType[T])
}
