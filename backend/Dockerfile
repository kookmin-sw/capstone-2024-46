FROM golang:1.20 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 go build -o server ./cmd/server/main.go

FROM gcr.io/distroless/static-debian12

WORKDIR /app

ENV PORT=8080
EXPOSE 8080

COPY docs/swagger-ui/ docs/swagger-ui/
COPY configs/ configs/

COPY --from=build /src/server ./

CMD ["/app/server"]
