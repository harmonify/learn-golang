//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

// Google wire use the param type in the providers
// to differentiate multiple params
func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}
