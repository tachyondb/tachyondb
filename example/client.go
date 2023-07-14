package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func saveToFile(filename string, data interface{}) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buffer.Bytes(), 0644)
}

func readFromFile(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	return dec.Decode(out)
}


func main() {
	filename := "users.gob"
	user := &User{FirstName: "John", LastName: "Doe"}
	users := []*User{user}

	// Save users to file
	if err := saveToFile(filename, users); err != nil {
		log.Fatalf("Failed to save: %v", err)
	}

	// Read users from file
	var loadedUsers []*User
	if err := readFromFile(filename, &loadedUsers); err != nil {
		log.Fatalf("Failed to load: %v", err)
	}

	fmt.Println("Loaded users:")
	for _, user := range loadedUsers {
		fmt.Printf("%+v\n", *user)
	}
}
