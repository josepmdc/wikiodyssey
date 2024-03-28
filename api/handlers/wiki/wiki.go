package wiki

import (
	"fmt"
	"net/http"

	"github.com/josepmdc/wikiodyssey/api/server"
	"github.com/josepmdc/wikiodyssey/api/services/wiki"
	"github.com/labstack/echo/v4"
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

func (h *Handler) GetArticlesTitles(e echo.Context, req server.GetArticlesTitlesParams) error {

	rawTitles, err := h.wikiService.GetTitles(req.Input)

	if err != nil {
		return fmt.Errorf("error getting titles: %w", err)
	}

	var titles []server.WikiPageObject

	for _, title := range rawTitles {
		titles = append(titles, server.WikiPageObject{
			Id:          title.Id,
			Title:       title.Title,
			Description: &title.Description,
			Key:         title.Key,
		})
	}

	return e.JSON(http.StatusOK, server.GetTitlesResponse{Titles: titles})
}

func (h *Handler) GetArticlesIsTitleInArticle(e echo.Context, req server.GetArticlesIsTitleInArticleParams) error {
	fmt.Printf("sourceTitle: %s | targetTitle: %s\n", req.SourceTitle, req.TargetTitle)
	isInArticle, err := h.wikiService.IsTitleInArticle(req.SourceTitle, req.TargetTitle)

	if err != nil {
		return fmt.Errorf("error checking if title is in article: %w", err)
	}

	return e.JSON(http.StatusOK, isInArticle)
}
