hello:
	echo "Hello World"

proto-generate:
	protoc ./proto/*.proto --proto_path=proto --go_out=plugins=grpc:. --go-grpc_out=.
