//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/josepmdc/wikiodyssey/api"
	wiki2 "github.com/josepmdc/wikiodyssey/api/handlers/wiki"
	"github.com/josepmdc/wikiodyssey/api/server"
	"github.com/josepmdc/wikiodyssey/api/services/wiki"
	"github.com/labstack/echo/v4"
)

func BuildApi() (*echo.Echo, error) {
	wire.Build(
		wire.Struct(new(api.Handlers), "*"),
		api.InitServer,
		wire.Bind(new(server.ServerInterface), new(*api.Handlers)),

		wiki.NewService,
		wiki2.NewHandler,
	)
	return &echo.Echo{}, nil
}
