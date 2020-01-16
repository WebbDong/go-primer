package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	code := base64.StdEncoding.EncodeToString([]byte("abcdefg2244234344ggbnnhhhjjjjjjggg"))
	fmt.Println(code)

	code2 := "YWJjZGVmZzIyNDQyMzQzNDRnZ2JubmhoaGpqampqamdnZw=="
	decodeString, err := base64.StdEncoding.DecodeString(code2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decodeString))
}
