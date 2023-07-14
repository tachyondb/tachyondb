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

func SaveToFile(filename string, data interface{}) (error) {
	dataArray := []interface{}{data}

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(dataArray)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromFile(filename string, data interface{}) (error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Decode the data
	var buffer2 bytes.Buffer
	buffer2.Write(fileData)

	decoder := gob.NewDecoder(&buffer2)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	startTime := time.Now()
	user := &User{
		FirstName: "John",
		LastName:  "Doe",
	}

	filename := "users.bin"

	SaveToFile(filename, user)

	newUsers := make([]User, 0)
	ReadFromFile(filename, &newUsers)

	log.Printf("Decoded User: %+v\n", newUsers)

	log.Print(time.Since(startTime))
}
