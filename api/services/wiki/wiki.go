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
