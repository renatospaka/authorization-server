## to list running services and their ports
ss -nltp

## protoc to generate stubs for gRPC
### message data
protoc --proto_path=./adapter/grpc/proto ./adapter/grpc/proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./adapter/grpc

### grpc data
protoc --proto_path=./adapter/grpc/proto ./adapter/grpc/proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./adapter/grpc
