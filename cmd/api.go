package main

import (
	"github.com/josepmdc/wikiodyssey/api/handlers"
	"github.com/josepmdc/wikiodyssey/api/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
	handlers.WikiHandler
}

func main() {
	api := Api{}
	e := echo.New()
	server.RegisterHandlers(e, &api)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
