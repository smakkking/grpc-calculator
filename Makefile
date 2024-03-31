.PHONY: gen-grpc
gen-grpc:
	protoc -I ./pkg/sdk/proto ./pkg/sdk/proto/*.proto \
	--go_out=./pkg/sdk/go --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/sdk/go --go_opt=paths=source_relative
