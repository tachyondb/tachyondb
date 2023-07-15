package driver

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
)

var filename = "driver.json" // TODO: this should be using the resource name later on

type Driver struct {}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Write(data interface{}) error {
	var fileContents []interface{}

	// Check if the file exists
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// File exists, read its contents
		err := d.Read(&fileContents)
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
	os.WriteFile(filename, b, 0644)

	return nil
}

func (d *Driver) BatchWrite(data interface{}) error {
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return errors.New("data must be a slice")
	}

	var fileContents []interface{}
	// Check if the file exists
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// File exists, read its contents
		err := d.Read(&fileContents)
		if err != nil {
			return err
		}
	}

	// Append new data to fileContents
	fileContents = append(fileContents, data.([]interface{})...)

	// Convert the slice to JSON
	b, err := json.Marshal(fileContents)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Write the JSON to the file
	os.WriteFile(filename, b, 0644)

	return nil
}

func (d *Driver) Read(data interface{}) error {
	// Read from file
	b, err := os.ReadFile(filename)
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
