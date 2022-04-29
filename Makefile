hello:
	echo "Hello World"

build-grpc:
	protoc --proto_path=proto --go_out=plugins=grpc:pb service.proto
