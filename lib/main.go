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
	https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch={next_node}&utf8=&format=json
3. Check if title of next node is linked in current node
	https://en.wikipedia.org/w/api.php?action=query&titles={current_node}&prop=links&pltitles={next_node}&format=json
*/

const randomApiURI = "https://en.wikipedia.org/w/api.php?action=query&format=json&list=random&rnnamespace=0"

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

	titles := make([]string, n)
	for i, entry := range randomResp.Query.Random {
		titles[i] = entry.Title
	}

	return titles, nil
}

func main() {
	GetRandomArticles(2)
}
