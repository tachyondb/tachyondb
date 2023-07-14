package main

import (
	"log"

	// "golang.org/x/net/context"
	// "google.golang.org/grpc"

	"github.com/tachyondb/tachyondb/users"
	"github.com/tachyondb/tachyondb/driver"
)

func main() {
	driver := driver.New()

	for i := 0; i < 10; i++ {
		user := &users.User{
			FirstName: "John",
			LastName: "Doe",
		}

		// TODO: change user to pointer
		if err := driver.Write("users", user); err != nil {
			log.Fatal(err)
		}
	}

	allUsers := &users.Users{}
	// user := &users.User{}

	// TODO: change user to pointer
	err := driver.Read("users", allUsers)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(allUsers)
}
