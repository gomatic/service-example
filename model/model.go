//go:generate protoc --proto_path=. -I/usr/local/include -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. model.proto
//go:generate protoc --proto_path=. -I/usr/local/include -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. model.proto
//go:generate protoc --proto_path=. -I/usr/local/include -I../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:.  model.proto

package model
