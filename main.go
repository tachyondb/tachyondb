package main

import (
	"log"
	"os"

	// "golang.org/x/net/context"
	// "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/tachyondb/tachyondb/users"
)

func SaveObject(T proto.Message) (error) {
	out, err := proto.Marshal(T)
	if err != nil {
		return err
	}

	if _, err := os.Stat("bin"); os.IsNotExist(err) {
		if err := os.Mkdir("bin", 0755); err != nil {
			return err
		}
	}

	if err := os.WriteFile("bin/users.bin", out, 0644); err != nil {
		return err
	}

	return nil
}

func GetObject(resource string) (proto.Message, error) {
	in, err := os.ReadFile("bin/" + resource + ".bin")
	if err != nil {
		return nil, err
	}

	allUsers := &users.Users{}

	if err := proto.Unmarshal(in, allUsers); err != nil {
		return nil, err
	}

	return allUsers, nil
}

func main() {
	allUsers := &users.Users{
		Users: []*users.User{},
	}

	user := &users.User{
		FirstName: "John",
		LastName: "Doe",
	}

	allUsers.Users = append(allUsers.Users, user)

	if err := SaveObject(allUsers); err != nil {
		log.Fatalln(err)
	}

	receivedUsers, err := GetObject("users")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(receivedUsers)
}