package main

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
	Name  string
	Value int
}

var filename = "json"



func main() {
	// Create an item and write it to the file
	for i := 0; i < 100_000; i++ {
		Write(Item{Name: "Item 1", Value: 1})
	}

	// Read all items from the file
	var itemsDecoded []interface{}
	Read(&itemsDecoded)

	log.Print("Items after write and read:", itemsDecoded)
}
