package driver

import (
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/tachyondb/tachyondb/users"
)

type Driver struct {}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Write(resource string, data proto.Message) error {
	allUsers := &users.Users{}

	filename := "bin/" + resource + ".bin"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		SaveObject(resource, allUsers)
		d.Read(resource, allUsers)
	}

	allUsers.Users = append(allUsers.Users, data.(*users.User))

	return SaveObject(resource, allUsers)
}

func (d *Driver) Read(resource string, data proto.Message) (error) {
	err := GetObject(resource, data)
	if err != nil {
		return err
	}

	return nil
}
