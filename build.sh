protoc --go_out=internal \
    --go_opt=paths=source_relative \
    --go-grpc_out=internal \
     --go-grpc_opt=paths=source_relative \
    -I=internal/ internal/base.grpc/agent.proto/proto.proto