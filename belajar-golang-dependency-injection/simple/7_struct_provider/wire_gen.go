// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

// Injectors from injector.go:

func InitializedFooBar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}
