# Private LLM Backend

## How to run

### Initialize

- Install [aws-cli](https://docs.aws.amazon.com/ko_kr/cli/latest/userguide/getting-started-install.html)
    ```sh
    brew install awscli
    ```
- Configure aws credential
    ```sh
    aws --profile blast configure
    ```

### Run

```sh
export AWS_PROFILE=blast
export ENV=local
go run cmd/server/main.go
```

## Code generator

### Dependency

- [protoc](https://grpc.io/docs/protoc-installation/) ^3.20
- [protoc-gen-openapi](https://github.com/google/gnostic/tree/main/cmd/protoc-gen-openapi) of gnostic
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [openapi-typescript-codegen](https://github.com/ferdikoomen/openapi-typescript-codegen)

### Install

```sh
brew install protobuf@3
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
npm install openapi-typescript-codegen -g
```

### Make

- `make` : _proto/*.proto_ 파일을 참고하여 _docs/openapi.yaml_ 을 생성하고, _internal/api/*.gen.go_ 에 type, server interface 코드와
  _dist/backend_ 에 client 코드를 생성합니다.

## Client SDK

서버에서 제공하는 API 를 쉽게 이용할 수 있는 client SDK 를 제공합니다.
[여기](dist) 를 참고해주세요.

## API

### Standard Error Response

200 이 아닌 모든 API 응답은 아래와 같은 형태로 리턴합니다.

```json
{
  "code": 3,
  "message": "bad request"
}
```

code 는 [gRPC status code](https://grpc.github.io/grpc/core/md_doc_statuscodes.html) 표준 spec 을 따르며, 아래는 상태 코드의 주요 목록과 그
의미입니다.

#### Status Code

| Code | Name                | Description                                                |
|------|---------------------|------------------------------------------------------------|
| 1    | Cancelled           | 작업이 클라이언트에 의해 취소되었음을 나타냅니다.                                |
| 2    | Unknown             | 알 수 없는 오류가 발생했음을 나타냅니다. 예상치 못한 예외가 발생한 경우에 사용됩니다.          |
| 3    | Invalid Argument    | 클라이언트가 잘못된 인자를 전달하여 요청이 거부되었음을 나타냅니다.                      |
| 4    | Deadline Exceeded   | 요청의 처리 시간이 너무 길어 지정된 마감 시간을 초과했음을 나타냅니다.                   |
| 5    | Not Found           | 지정된 리소스를 찾을 수 없음을 나타냅니다.                                   |
| 6    | Already Exists      | 생성하려는 리소스가 이미 존재함을 나타냅니다.                                  |
| 7    | Permission Denied   | 적절한 권한 없이 리소스에 접근하려고 했음을 나타냅니다.                            |
| 8    | Resource Exhausted  | 리소스 할당량 초과 등으로 요청을 처리할 수 없음을 나타냅니다.                        |
| 9    | Failed Precondition | 시스템 상태가 요청을 처리하기에 적합하지 않음을 나타냅니다.                          |
| 10   | Aborted             | 연산이 중단되었음을 나타냅니다. 일반적으로 동시성 문제나 데이터 무결성 충돌 때문입니다.          |
| 11   | Out of Range        | 연산이 허용된 범위를 벗어났음을 나타냅니다.                                   |
| 12   | Unimplemented       | 요청된 연산이 서버에서 구현되지 않았음을 나타냅니다.                              |
| 13   | Internal            | 내부 오류가 발생했음을 나타냅니다. 일반적으로 서버 측 문제입니다.                      |
| 14   | Unavailable         | 서비스가 일시적으로 사용할 수 없음을 나타냅니다. 일시적인 오버로드나 유지보수로 인한 것일 수 있습니다. |
| 15   | Data Loss           | 데이터 손실이 발생했음을 나타냅니다.                                       |
| 16   | Unauthenticated     | 요청이 인증되지 않았음을 나타냅니다. 유효한 인증 자격증명이 제공되지 않았을 때 발생합니다.        |
