package driver

import (
	"google.golang.org/protobuf/proto"
)

type Driver struct {
}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Write(resource string, data proto.Message) error {
	return SaveObject(resource, data)
}

func (d *Driver) Read(resource string, data proto.Message) (proto.Message, error) {
	return GetObject(resource, data)
}
