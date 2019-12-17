package main

import (
	"fmt"
)

const apiKey = ""

func main() {

	items, err := getItems(apiKey)
	if err != nil {
		fmt.Printf("Get Items error\n") 
	} else {
		for _, i := range items.Data {
			fmt.Printf(i.Name)
			fmt.Printf("\n")
		}
	}
}
