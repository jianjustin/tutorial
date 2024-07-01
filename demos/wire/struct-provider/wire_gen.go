// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitializeFooBar() FooBar {
	foo := ProvideFoo()
	bar := ProvideBar()
	fooBar := FooBar{
		MyFoo: foo,
		MyBar: bar,
	}
	return fooBar
}