package main

import "github.com/josepmdc/wikiodyssey/api/wire"

func main() {
	_, err := wire.BuildApi()
	if err != nil {
		panic(err)
	}
}
