protoc -I ./proto \
   --go_out ./stubs --go_opt paths=source_relative \
   --go-grpc_out ./stubs --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./stubs --grpc-gateway_opt paths=source_relative \
   ./proto/*/v1/*.proto

protoc -I ./proto \
  --openapiv2_out ./api \
  --openapiv2_opt logtostderr=true \
  --openapiv2_opt use_go_templates=true \
  ./proto/*/v1/*_service.proto