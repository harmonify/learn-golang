//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}
