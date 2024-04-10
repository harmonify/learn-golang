//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}
