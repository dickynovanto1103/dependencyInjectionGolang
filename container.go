//+build wireinject

package main

import (
	"github.com/dickynovanto1103/dependencyInjectionGolang/httpClient"
	"github.com/dickynovanto1103/dependencyInjectionGolang/logger"
	"github.com/dickynovanto1103/dependencyInjectionGolang/serviceGroup"
	"github.com/google/wire"
)

func CreateServiceGroup() *ServiceGroup {
	panic(wire.Build(
		httpClient.NewHttpClient,
		serviceGroup.NewServiceGroup,
		wire.Struct(new(logger.Logger)),
	))
}
