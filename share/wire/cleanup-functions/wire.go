// +build wireinject

package main

import "github.com/google/wire"

func InitializeFileReader(path string) (reader *FileReader, cleanup func(), err error) {
	//func InitializeFileReader(path string) (*FileReader, func(), error) {
	wire.Build(NewFileReader)
	return
}
