//go:build wireinject
// +build wireinject

package infra

import "github.com/google/wire"

func BuildServer() *ServerBuilder {
	wire.Build(AppSet)

	return &ServerBuilder{}
}
