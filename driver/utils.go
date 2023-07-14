package driver

import (
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/tachyondb/tachyondb/users"
)

func SaveObject(resource string, T proto.Message) (error) {
	out, err := proto.Marshal(T)
	if err != nil {
		return err
	}

	if _, err := os.Stat("bin"); os.IsNotExist(err) {
		if err != nil {
			return err
		}

		if err := os.Mkdir("bin", 0755); err != nil {
			return err
		}
	}

	if err := os.WriteFile("bin/" + resource + ".bin", out, 0644); err != nil {
		return err
	}

	return nil
}

func GetObject(resource string, data proto.Message) (proto.Message, error) {
	in, err := os.ReadFile("bin/" + resource + ".bin")
	if err != nil {
		return nil, err
	}

	allUsers := &users.Users{}


	if err := proto.Unmarshal(in, data); err != nil {
		return nil, err
	}

	return allUsers, nil
}