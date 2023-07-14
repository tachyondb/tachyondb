package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Item struct {
	Name  string
	Value int
}

var filename = "data.bin"

func Write(item interface{}) error {
	// Open the file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new encoder writing to the file
	enc := gob.NewEncoder(file)

	// Wrap the item in a slice and encode it
	err = enc.Encode([]interface{}{item})
	if err != nil {
		return err
	}

	return nil
}


func Read(items interface{}) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a decoder
	dec := gob.NewDecoder(file)

	// Decode (read) the items into the provided slice
	sliceValue := reflect.ValueOf(items).Elem()
	for {
		itemType := sliceValue.Type().Elem()
		item := reflect.New(itemType)

		err := dec.Decode(item.Interface())
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return err
			}
		}

		sliceValue.Set(reflect.Append(sliceValue, item.Elem()))
	}

	return nil
}


func main() {
	gob.Register(Item{})

	items := []Item{
		{"Item 1", 1},
		{"Item 2", 2},
		{"Item 3", 3},
	}

	for _, item := range items {
		err := Write(item)
		if err != nil {
			log.Fatal(err)
		}
	}

	var newItems []Item
	err := Read(&newItems)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range newItems {
		fmt.Printf("%s: %d\n", item.Name, item.Value)
	}
}
