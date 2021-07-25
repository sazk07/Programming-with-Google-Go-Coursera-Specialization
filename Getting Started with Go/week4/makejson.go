package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Address: ")
	fmt.Scan(&address)

	var nameAddress map[string]string
	nameAddress = make(map[string]string)
	nameAddress["name"] = name
	nameAddress["address"] = address

	i, _ := json.Marshal(nameAddress)

	fmt.Println(string(i))
}
