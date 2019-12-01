//+build wireinject

package main

import (
	"github.com/google/wire"
)

func CreateServiceGroup() *ServiceGroup {
	panic(wire.Build(
		NewHttpClient,
		NewServiceGroup,
		wire.Struct(new(Logger)),
	))
}
