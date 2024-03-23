package handlers

import (
	"fmt"
	"github.com/josepmdc/wikiodyssey/api/server"
	"github.com/josepmdc/wikiodyssey/lib/wiki"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WikiHandler struct{}

func (*WikiHandler) GetArticlesRandom(e echo.Context, req server.GetArticlesRandomParams) error {
	limit := uint(1)
	if req.Limit != nil {
		limit = *req.Limit
	}

	articles, err := wiki.GetRandomArticles(limit)
	if err != nil {
		return fmt.Errorf("failed to get random articles: %w", err)
	}

	return e.JSON(http.StatusOK, server.RandomArticlesResponse{Articles: articles})
}
