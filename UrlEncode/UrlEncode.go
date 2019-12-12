package main

import (
	"fmt"
	"net/url"
)

func main() {
	values := url.Values{}
	values.Add("data", "abcd+23432++fdsa")
	encode := values.Encode()
	fmt.Println(encode)

	query, err := url.ParseQuery(encode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(query.Get("data"))
}
