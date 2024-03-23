package api

import (
	"github.com/josepmdc/wikiodyssey/api/handlers/wiki"
	"github.com/josepmdc/wikiodyssey/api/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct {
	*wiki.Handler
}

func InitServer(handlers server.ServerInterface /* TODO: Inject config here */) *echo.Echo {
	e := echo.New()
	server.RegisterHandlers(e, handlers)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
	return e
}
