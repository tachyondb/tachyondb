package main

import (
	"log"

	driverInterface "github.com/tachyondb/tachyondb/driver"
)

type Item struct {
	Name  string
	Value int
}

var filename = "json"



func main() {
	driver := driverInterface.New()

	// Create an item and write it to the file
	driver.Write(Item{Name: "Item 1", Value: 1})

	arrayOfItems := []*Item{}

	for i := 0; i < 100_000; i++ {
		arrayOfItems = append(arrayOfItems, &Item{Name: "Item 1", Value: 1})
	}

	driver.BatchWrite(arrayOfItems)

	// Read all items from the file
	var itemsDecoded []interface{}
	driver.Read(&itemsDecoded)

	log.Print("Items after write and read:", itemsDecoded)
}
