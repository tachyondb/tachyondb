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

func Write(data interface{}) error {
	var fileContents []interface{}

	// Check if bson.bin file exists
	if _, err := os.Stat("bson.bin"); !os.IsNotExist(err) {
		// File exists, read its contents
		err := Read(&fileContents)
		if err != nil {
			log.Fatal("read error:", err)
		}
	}

	// Append new data to fileContents
	fileContents = append(fileContents, data)

	// Convert the slice to JSON
	b, err := json.Marshal(fileContents)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Write the JSON to the file
	os.WriteFile("bson.bin", b, 0644)

	return nil
}

func Read(data interface{}) error {
	// Read from file
	b, err := os.ReadFile("bson.bin")
	if err != nil {
		return err
	}

	// Unmarshal the JSON into the data
	err = json.Unmarshal(b, data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Create an item and write it to the file
	item := Item{
		Name:  "Item 1",
		Value: 1,
	}
	Write(item)

	// Write some more items to the file
	Write(Item{Name: "Item 2", Value: 2})
	Write(Item{Name: "Item 3", Value: 3})

	// Read all items from the file
	var itemsDecoded []interface{}
	Read(&itemsDecoded)

	log.Print("Items after write and read:", itemsDecoded)
}
