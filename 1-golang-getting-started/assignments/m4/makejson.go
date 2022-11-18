package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string
	fmt.Print("enter a name: ")
	fmt.Scan(&name)
	fmt.Print("enter an address: ")
	fmt.Scan(&address)

	b, err := json.Marshal(map[string]string{
		"name":    name,
		"address": address,
	})
	if err != nil {
		fmt.Println("failed to marshal map to json object")
		os.Exit(1)
		return
	}
	fmt.Println("json object:", string(b))
}
