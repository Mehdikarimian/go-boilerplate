FROM golang:alpine as builder

WORKDIR /go/app

COPY src ./src
COPY docs ./docs
COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN apk update && apk add --no-cache git
RUN go mod download

RUN go build -o ./src/build ./src

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/app/src/build .

EXPOSE $PORT

ENTRYPOINT ["./build"]