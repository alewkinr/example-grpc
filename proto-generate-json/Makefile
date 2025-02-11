MODULE_NAME = github.com/alewkinr/example-grpc/proto-generate-json/sdk/go

.PHONY: lint
lint: format
	buf lint --config buf.yaml --timeout 5m0s

.PHONY: format
format:
	buf format proto -w

.PHONY: generate
generate: format lint
	buf generate --template buf.gen.yaml
	cd ./sdk/go && go mod init $(MODULE_NAME) || true
	go mod tidy

.PHONY:
clean: clean-sdk

.PHONY: clean-sdk
clean-sdk:
	rm -rf ./sdk/go

