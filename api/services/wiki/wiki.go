package wiki

import (
	"fmt"

	"github.com/josepmdc/wikiodyssey/lib/wiki"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (svc *Service) GetRandomArticles(limit uint) ([]string, error) {
	articles, err := wiki.GetRandomArticles(limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get random articles: %w", err)
	}
	return articles, nil
}

func (svc *Service) GetTitles(input string) ([]*wiki.WikiPageObject, error) {
	pages, err := wiki.GetTitles(input)
	if err != nil {
		return nil, fmt.Errorf("failed getting titles with input  '%s': %w", input, err)
	}
	return pages, nil
}

func (svc *Service) IsTitleInArticle(sourceTitle string, targetTitle string) (string, error) {
	nextTitle, err := wiki.IsTitleInArticle(sourceTitle, targetTitle)

	if err != nil {
		return "", fmt.Errorf("failed checking if target title '%s' is in '%s' article: %w", sourceTitle, targetTitle, err)
	}

	return nextTitle, nil
}
