package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"time"
)

type User struct {
	FirstName string
	LastName string
}

func SaveToFile(filename string, data interface{}) error {
	buffer := &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)

	if err := enc.Encode(data); err != nil {
		return err
	}

	err := os.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromFile(filename string, data interface{}) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(fileData)
	dec := gob.NewDecoder(buffer)

	if err := dec.Decode(data); err != nil {
		return err
	}

	return nil
}


func main() {
	startTime := time.Now()

	var data interface{}
	data = &User{
		FirstName: "John",
		LastName:  "Doe",
	}

	filename := "users.bin"

	if err := SaveToFile(filename, data); err != nil {
		log.Fatal(err)
	}

	data = &User{}
	if err := ReadFromFile(filename, data); err != nil {
		log.Fatal(err)
	}

	log.Printf("Decoded User: %+v\n", data)

	log.Print(time.Since(startTime))
}
