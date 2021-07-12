FROM golang:1.14.2-alpine as builder

ENV GIN_MODE=release

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o app .

# Run image
FROM alpine:latest

ENV GIN_MODE=release

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
