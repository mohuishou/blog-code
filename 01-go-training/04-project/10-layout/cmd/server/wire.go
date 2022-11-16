//+build wireinject

package main

import (
	"github.com/google/wire"
)

// NewServices NewServices
func NewServices() (*services, error) {
	panic(wire.Build(
		wire.Struct(new(services), "*"),
		set,
	))
}
