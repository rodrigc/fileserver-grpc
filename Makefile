
.PHONY: generate server client test clean

generate:
	@if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc --go_out=. --go-grpc_out=.  proto/file.proto

server:
	go build -o out/server ./cmd/server.go

client:
	go build -o out/client ./client/client.go


test:
	@if ! which grpcurl > /dev/null; then \
		echo "error: grpcurl not installed" >&2; \
		exit 1; \
	fi
	grpcurl -import-path ./proto -proto file.proto list
	grpcurl -plaintext -proto proto/file.proto -import-path . -d '{"file_name": "server"}' localhost:8081 fileservice.FileService/Exists
	grpcurl -plaintext -proto proto/file.proto -import-path . -d '{"file_name": "notexists.txt"}' localhost:8081 fileservice.FileService/Exists || true
	grpcurl -plaintext -proto proto/file.proto -import-path . -d '{"file_name": "server"}' localhost:8081 fileservice.FileService/GetFileMetaData
	grpcurl -plaintext -proto proto/file.proto -import-path . -d '{"file_name": "notexists.txt"}' localhost:8081 fileservice.FileService/GetFileMetaData || true

clean:
	rm -fr out
