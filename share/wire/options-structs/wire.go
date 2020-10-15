// +build wireinject

package main

import (
	"context"
	"io"

	"github.com/google/wire"
)

func InitializeMyStruct(ctx *context.Context, msgs []Message, reader io.Reader, writer io.Writer) (MyStruct, error) {
	wire.Build(Set)
	return MyStruct{}, nil
}
