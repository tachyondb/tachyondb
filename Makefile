start:
	go run .

proto:
	mkdir -p users
	protoc -I protos/ --go_out=plugins=grpc:users --go_opt=paths=source_relative --go_opt=Muser.proto=github.com/tachyondb/tachyondb/users user.proto
