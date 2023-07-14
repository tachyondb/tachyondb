package main

import (
	"errors"
	"log"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

type Item struct {
	Name  string
	Value int
}

func Write(data interface{}) {
	// Encode the item
	b, err := bson.Marshal(data)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// save to file
	os.WriteFile("bson.bin", b, 0644)
}

func Read(data interface{}) error {
	// Read from file
	b, err := os.ReadFile("bson.bin")
	if err != nil {
		return err
	}

	// Ensure data is a pointer
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr {
		return errors.New("data must be a pointer")
	}

	// Check for nil pointer before dereferencing
	if rv.IsNil() {
		return errors.New("data cannot be a nil pointer")
	}

	// Create a new instance of the type pointed to by data
	// Note: This requires that the underlying type is publicly accessible
	newData := reflect.New(rv.Elem().Type())

	// Unmarshal BSON into the newly created data
	err = bson.Unmarshal(b, newData.Interface())
	if err != nil {
		return err
	}

	// Now assign the value back to data
	rv.Elem().Set(newData.Elem())

	log.Print(data, reflect.TypeOf(data))

	return nil
}



func main() {
	// Create an item
	item := &Item{
		Name:  "Item 1",
		Value: 1,
	}
	Write(item)

	// Decode the item
	itemDecoded := &Item{}
	Read(itemDecoded)

	log.Print("item after write and read:", itemDecoded)
}
