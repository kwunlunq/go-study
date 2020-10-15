package main

import (
	"fmt"

	"github.com/google/wire"
)

type Foo int
type Bar int

func NewFoo() Foo {
	return 1
}

func NewBar() Bar {
	return 2
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

var Set = wire.NewSet(
	NewFoo,
	NewBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"),
	// == wire.Struct(new(FooBar), "*")), inject all fields
)

func main() {
	fooBar := InitializeFooBar()
	fmt.Printf(" %v \n", fooBar)
}
