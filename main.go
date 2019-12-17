package main

import (
	"fmt"
)

const apiKey = ""
const tenantUrl = ""

func main() {

	items, err := getItem(apiKey, tenantUrl, "Bugs")
	if err != nil {
		fmt.Printf("Get Item error\n") 
	} else {
		for _, i := range items.Data {
			fmt.Printf("reloading " + i.Name + "\n")
			_, err := reload(apiKey, tenantUrl, i.ResourceId)
			if err != nil {
				fmt.Printf("Reload error\n") 
			}
		}
	}
}
