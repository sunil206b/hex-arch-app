protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto \
 internal/adapters/framework/left/grpc/proto/numbers_msg.proto

protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/grpc \
--proto_path=internal/adapters/framework/left/grpc/proto \
internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto