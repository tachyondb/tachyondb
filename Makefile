start:
	go run .

proto:
	protoc -I protos/ --go_out=plugins=grpc:protos --go_opt=paths=source_relative --go_opt=Morder.proto=github.com/tachyondb/tachyondb/users user.proto
