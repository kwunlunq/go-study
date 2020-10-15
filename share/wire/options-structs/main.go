package main

import (
	"context"
	"io"

	"github.com/google/wire"
)

type MyStruct struct{}

func NewMyStruct(ctx *context.Context, opt Option) (MyStruct, error) {
	return MyStruct{}, nil
}

type Option struct {
	Messages []Message
	Writer   io.Writer
	Reader   io.Reader
}

type Message string

var Set = wire.NewSet(NewMyStruct, wire.Struct(new(Option), "*"))
