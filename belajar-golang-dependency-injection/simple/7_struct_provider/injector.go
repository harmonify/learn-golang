//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}
