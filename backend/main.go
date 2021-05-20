package main

import (
	"net/url"
)

func main() {

	query := make(map[string]string)
	query["a"] = "b"
	query["c"] = "d"

	params := url.Values{}
	for key, value := range query {
		params.Add(key, value)
	}

}
