package main

import (
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

func Read(data interface{}) ([]byte, error) {
	// Read from file
	b, err := os.ReadFile("bson.bin")
	if err != nil {
		log.Fatal("read error:", err)
	}

	dataAfterTypeCheck := reflect.TypeOf(data)
	dataType := reflect.TypeOf(data).Kind()
	dataSome := reflect.New(dataAfterTypeCheck)
	log.Print(dataAfterTypeCheck, dataType, dataSome)

	err = bson.Unmarshal(b, &data)

	if err != nil {
		log.Fatal("decode error:", err)
	}

	log.Print(data, reflect.TypeOf(data))

	return b, nil
}

func main() {
	// Create an item
	item := &Item{
		Name:  "Item 1",
		Value: 1,
	}
	log.Print("item before write:", item)
	Write(item)

	// Decode the item
	var itemDecoded Item
	b, err := Read(itemDecoded)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	err = bson.Unmarshal(b, &itemDecoded)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	log.Print("item after write and read:", itemDecoded)
}
