FROM golang:1.23.5 AS build-stage

WORKDIR /merch-shop

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY internal/ ./internal
COPY migrations/ ./migrations

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest

WORKDIR /merch-shop

COPY --from=build-stage /merch-shop/main .

EXPOSE 8080

CMD ["./main"]
