package main

import (
	"log"
	"os"

	"github.com/tachyondb/tachyondb/users"
	"google.golang.org/protobuf/proto"
)

func main() {
	fileContents, err := os.ReadFile("json")
	if err != nil {
		log.Fatal(err)
	}

	fileContentsStr := string(fileContents)
	str := &users.JSONString{
		Json: fileContentsStr,
	}

	b, err := proto.Marshal(str)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("proto.bin", b, 0644)

	// read the file
	in, err := os.ReadFile("proto.bin")

	if err != nil {
		log.Fatal(err)
	}

	str2 := &users.JSONString{}

	if err := proto.Unmarshal(in, str2); err != nil {
		log.Fatal(err)
	}

	log.Println(str2)


	c := []byte(str2.Json)
	os.WriteFile("test.json", c, 0644)
}
