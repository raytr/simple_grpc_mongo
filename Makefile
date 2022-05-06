hello:
	echo "Hello World"

proto-generate:
	protoc -I ./proto \
	   --go_out ./gen --go_opt paths=source_relative \
	   --go-grpc_out ./gen --go-grpc_opt paths=source_relative \
	   ./proto/*.proto

proto-generate-2:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	protoc -I . --grpc-gateway_out ./gen/ \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        proto/*.proto

gen-google-api:
	mkdir -p google/api
	curl https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto > proto/google/api/annotations.proto
	curl https://github.com/googleapis/googleapis/blob/master/google/api/http.proto > proto/google/api/http.proto