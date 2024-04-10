//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
