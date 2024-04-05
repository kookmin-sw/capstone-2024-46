# Makefile

.PHONY: all generate-docs generate-server generate-client generate-openai-client

all: generate-docs generate-server generate-client

generate-docs:
	@echo ''
	@protoc proto/*.proto -I=proto/ -I=proto/third_party --openapi_out=docs --openapi_opt=enum_type=string && \
		go run cmd/update_enum_prefix/main.go docs/openapi.yaml

generate-server:
	@oapi-codegen --generate=echo-server,strict-server,embedded-spec --package=api -o internal/api/server.gen.go  docs/openapi.yaml
	@oapi-codegen --generate=models --package=api -o internal/api/types.gen.go  docs/openapi.yaml
	@go run private-llm-backend/cmd/type_keys_generator
	@go fmt private-llm-backend/internal/api

generate-client:
	@protoc proto/*.proto -I=proto/ -I=proto/third_party --openapi_out=dist --openapi_opt=default_response=false --openapi_opt=enum_type=string && \
		go run cmd/update_enum_prefix/main.go dist/openapi.yaml && \
		npx openapi-typescript-codegen -i dist/openapi.yaml -o dist/backend --name BackendClient --client fetch && \
		rm dist/openapi.yaml && \
		cp -r dist/.template/* dist/backend/

generate-openai-client:
	@protoc proto/client/openai/*.proto -I=proto/ -I=proto/third_party --openapi_opt=default_response=false --openapi_opt=naming=proto --openapi_opt=enum_type=string --openapi_out=proto/client/openai/ && \
	oapi-codegen --generate=client,models --package=openai -o pkg/client/openai/openai.gen.go proto/client/openai/openapi.yaml && \
	rm proto/client/openai/*.yaml
