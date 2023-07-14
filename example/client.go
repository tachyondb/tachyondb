package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
)

type Item struct {
	Name  string
	Value int
}

var filename = "data.bin"

func Write(item interface{}) error {
		// Create a buffer to store encoded data
		var buf bytes.Buffer

		// Create a new encoder writing to the buffer
		enc := gob.NewEncoder(&buf)

		// Encode (write) the items one by one
		err := enc.Encode(&item)
		if err != nil {
			return err
		}

		// Write the encoded data to a file
		err = ioutil.WriteFile(filename, buf.Bytes(), 0644)

		if err != nil {
			return err
		}

		return nil
}

func Read() {

}

func main() {
	items := []Item{
		{"Item 1", 1},
		{"Item 2", 2},
		{"Item 3", 3},
	}
	Write(items[0])
	// Now let's read the data back and decode it
	data, err := ioutil.ReadFile("data.gob")
	if err != nil {
		log.Fatal("read file:", err)
	}

	// Create a decoder
	dec := gob.NewDecoder(bytes.NewBuffer(data))

	// Decode (read) the items into a new slice
	var decodedItems []Item
	for {
		var item Item
		err := dec.Decode(&item)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Fatal("decode:", err)
			}
		}
		decodedItems = append(decodedItems, item)
	}

	// Print the decoded items
	for _, item := range decodedItems {
		fmt.Printf("%s: %d\n", item.Name, item.Value)
	}
}
