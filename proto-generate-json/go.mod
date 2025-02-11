module github.com/alewkinr/example-grpc

go 1.23.2

replace github.com/alewkinr/example-grpc/proto-generate-json/sdk/go => ./sdk/go

require github.com/alewkinr/example-grpc/proto-generate-json/sdk/go v0.0.0-00010101000000-000000000000

require google.golang.org/protobuf v1.36.5 // indirect
