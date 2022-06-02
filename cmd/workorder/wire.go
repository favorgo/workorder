//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pipperman/workorder/internal/biz"
	"github.com/pipperman/workorder/internal/conf"
	"github.com/pipperman/workorder/internal/data"
	"github.com/pipperman/workorder/internal/server"
	"github.com/pipperman/workorder/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	//"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
