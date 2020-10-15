// +build wireinject

package main

import "github.com/google/wire"

func InitializeFooBar() FooBar {
	wire.Build(Set)
	return FooBar{}
}
