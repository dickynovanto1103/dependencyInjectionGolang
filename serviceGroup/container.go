//+build wireinject

package serviceGroup

import (
	"github.com/dickynovanto1103/dependencyInjectionGolang/httpClient"
	"github.com/dickynovanto1103/dependencyInjectionGolang/logger"
	"github.com/google/wire"
)

func CreateServiceGroup() *ServiceGroup {
	panic(wire.Build(
		httpClient.NewHttpClient,
		NewServiceGroup,
		wire.Struct(new(logger.Logger)),
	))
}
