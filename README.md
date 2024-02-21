

### Generate protos and stubs

// To get protoc to work
export PATH="$PATH:$(go env GOPATH)/bin"

protoc \
    --go_out=internal/adapters/framework/left_driver/grpc \
    --proto_path=internal/adapters/framework/left_driver/grpc/proto \
    internal/adapters/framework/left_driver/grpc/proto/number_msg.proto

protoc \
    --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left_driver/grpc \
    --proto_path=internal/adapters/framework/left_driver/grpc/proto \
    internal/adapters/framework/left_driver/grpc/proto/arithmetic_svc.proto


