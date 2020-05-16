package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
	   Write a program which prompts the user to first enter a name, and then enter an address.
	   Your program should create a map and add the name and address to the map using the keys “name” and
	    “address”, respectively. Your program should use Marshal() to create a JSON object from the map,
	    and then your program should print the JSON object.
	*/
	var name1 string
	var address1 string

	fmt.Println("Please enter your name: ")
	fmt.Scanln(&name1)
	fmt.Println("Please enter your address: ")
	fmt.Scanln(&address1)

	map1 := make(map[string]string)

	map1["name"] = name1
	map1["address"] = address1

	barr, _ := json.Marshal(map1)
	fmt.Println(string(barr))
}
