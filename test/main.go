package main

import (
	"fmt"
	"github.com/jimareed/qlikapi"
)

const apiKey = ""
const tenantUrl = ""

func main() {

	items, err := qlikapi.GetItem(apiKey, tenantUrl, "Bugs")
	if err != nil {
		fmt.Printf("Get Item error\n") 
	} else {
		for _, i := range items.Data {
			fmt.Printf("reloading " + i.Name + "\n")
			_, err := qlikapi.Reload(apiKey, tenantUrl, i.ResourceId)
			if err != nil {
				fmt.Printf("Reload error\n") 
			}
		}
	}
}
