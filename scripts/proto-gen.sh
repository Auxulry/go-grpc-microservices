protoc -I ./proto \
   --go_out ./stubs --go_opt paths=source_relative \
   --go-grpc_out ./stubs --go-grpc_opt paths=source_relative \
   ./proto/*/v1/*.proto