package main

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// Convert to binary and write to file
	binFile, err := os.Create("person.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer binFile.Close()

	enc := gob.NewEncoder(binFile)
	err = enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON and write to file
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("person.json", jsonData, 0644); err != nil {
		log.Fatal(err)
	}
}
