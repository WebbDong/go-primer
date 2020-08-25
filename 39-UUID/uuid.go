package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

// go get -u github.com/satori/go.uuid
func main() {
	value, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value.String())

	fromString, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fromString)
}
