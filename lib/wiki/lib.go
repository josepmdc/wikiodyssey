package wiki

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
1. Get two random articles
	https://en.wikipedia.org/w/api.php?action=query&format=json&list=random&rnlimit=2&rnnamespace=0
2. Search for titles based on input:
	https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch={next_node}&utf8=&format=json&rnnamespace=0
3. Check if title of next node is linked in current node (Current_node has to be exact title, caps included)
	https://en.wikipedia.org/w/api.php?action=query&titles={current_node}&prop=links&pltitles={next_node}&format=json
*/

const randomApiURI = "https://en.wikipedia.org/w/api.php?action=query&format=json&list=random&rnnamespace=0"
const articleLinksURI = "https://en.wikipedia.org/w/api.php?action=query&titles=%s&prop=links&pltitles=%s&format=json&pllimit=max&plnamespace=0"
const searchTitlesURI = "https://en.wikipedia.org/w/rest.php/v1/search/title?q=%s&limit=10"

type RandomQuery struct {
	Query *WikiRandom `json:"query"`
}

type WikiRandom struct {
	Random []*RandomEntry `json:"random"`
}

type RandomEntry struct {
	Title string `json:"title"`
}

func GetRandomArticles(n uint) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s&rnlimit=%d", randomApiURI, n))
	if err != nil {
		return nil, fmt.Errorf("error fetching random articles: %w", err)
	}

	var randomResp RandomQuery
	err = json.NewDecoder(resp.Body).Decode(&randomResp)
	if err != nil {
		return nil, fmt.Errorf("failed to get random article")
	}

	titles := make([]string, n)
	for i, entry := range randomResp.Query.Random {
		titles[i] = entry.Title
	}

	return titles, nil
}

type SearchTitleQuery struct {
	Pages []*WikiPageObject `json:"pages"`
}

type WikiPageObject struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetTitles(input string) ([]*WikiPageObject, error) {
	input = strings.ReplaceAll(input, " ", "_")
	resp, err := http.Get(fmt.Sprintf(searchTitlesURI, input))
	if err != nil {
		return nil, fmt.Errorf("error fetching article title: %w", err)
	}

	var searchResp SearchTitleQuery
	err = json.NewDecoder(resp.Body).Decode(&searchResp)

	if err != nil {
		return nil, fmt.Errorf("error decoding titles: %w", err)
	}

	return searchResp.Pages, nil
}

type CheckerQuery struct {
	Query struct {
		Pages map[string]struct {
			Links []*CheckerEntry `json:"links"`
		} `json:"pages"`
	} `json:"query"`
}

type CheckerEntry struct {
	Title string `json:"title"`
}

func IsTitleInArticle(sourceTitle string, targetTitle string) (string, error) {
	sourceTitle = strings.ReplaceAll(sourceTitle, " ", "_")
	originalTargetTitle := targetTitle
	targetTitle = strings.ReplaceAll(targetTitle, " ", "_")
	resp, err := http.Get(fmt.Sprintf(articleLinksURI, sourceTitle, targetTitle))

	if err != nil {
		return "", fmt.Errorf("error fetching random articles: %w", err)
	}

	var checkerQuery CheckerQuery
	err = json.NewDecoder(resp.Body).Decode(&checkerQuery)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON: %w", err)
	}

	links := make([]*CheckerEntry, 1)

	for _, value := range checkerQuery.Query.Pages {
		if len(value.Links) == 1 {
			return value.Links[0].Title, nil
		} else {
			links = value.Links
		}

		break
	}

	for _, link := range links {
		if link.Title == originalTargetTitle {
			return link.Title, nil
		}
	}

	return "", fmt.Errorf("no matching titles found")
}
