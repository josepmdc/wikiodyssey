package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
1. Get two random articles
	https://en.wikipedia.org/w/api.php?action=query&format=json&list=random&rnlimit=2&rnnamespace=0
2. Search for titles based on input:
	https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch={next_node}&utf8=&format=json&rnnamespace=0
3. Check if title of next node is linked in current node
	https://en.wikipedia.org/w/api.php?action=query&titles={current_node}&prop=links&pltitles={next_node}&format=json
*/

const randomApiURI = "https://en.wikipedia.org/w/api.php?action=query&format=json&list=random&rnnamespace=0"
const searchTitlesURI = "https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=%s&utf8=&format=json"

type RandomQuery struct {
	Query *WikiRandom `json:"query"`
}

type WikiRandom struct {
	Random []*RandomEntry `json:"random"`
}

type RandomEntry struct {
	Title string `json:"title"`
}

type SearchTitleQuery struct {
	Query *WikiSearch `json:"query"`
}

type WikiSearch struct {
	SearchInfo struct {
		TotalHits uint `json:"totalhits"`
	} `json:"searchinfo"`

	Search []*SearchEntry `json:"search"`
}

type SearchEntry struct {
	Title string `json:"title"`
}

// type SearchInformation struct {
// 	TotalHits uint `json:"totalhits"`
// }

func GetRandomArticles(n uint) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s&rnlimit=%d", randomApiURI, n))
	if err != nil {
		return nil, fmt.Errorf("error fetching random articles: %w", err)
	}

	var randomResp RandomQuery
	err = json.NewDecoder(resp.Body).Decode(&randomResp)

	titles := make([]string, n)
	for i, entry := range randomResp.Query.Random {
		titles[i] = entry.Title
	}

	return titles, nil
}

func GetTitles(input string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(searchTitlesURI, input)) //replace spaces?

	if err != nil {
		return "", fmt.Errorf("error fetching random articles: %w", err)
	}

	var searchResp SearchTitleQuery
	err = json.NewDecoder(resp.Body).Decode(&searchResp)

	if searchResp.Query.SearchInfo.TotalHits <= 0 {
		return "", fmt.Errorf("error fetching random articles: no results found")
	}

	title := searchResp.Query.Search[0].Title

	return title, nil
}

func main() {
	title, _ := GetTitles("einstein")

	fmt.Println(title)

}
