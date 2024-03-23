package wiki

import (
	"fmt"
	"github.com/josepmdc/wikiodyssey/api/server"
	"github.com/josepmdc/wikiodyssey/api/services/wiki"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	wikiService *wiki.Service
}

func NewHandler(wikiService *wiki.Service) *Handler {
	return &Handler{
		wikiService: wikiService,
	}
}

func (h *Handler) GetArticlesRandom(e echo.Context, req server.GetArticlesRandomParams) error {
	limit := uint(1)
	if req.Limit != nil {
		limit = *req.Limit
	}

	articles, err := h.wikiService.GetRandomArticles(limit)
	if err != nil {
		return fmt.Errorf("error getting random articles: %w", err)
	}

	return e.JSON(http.StatusOK, server.RandomArticlesResponse{Articles: articles})
}
