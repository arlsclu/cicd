FROM golang:1.19.4-buster as builder

WORKDIR /app

COPY go.mod .
RUN go mod tidy
COPY *.go /app/
RUN go build -o hello

FROM alpine:latest

COPY --from=builder /app/hello /hello

ENTRYPOINT [ "/hello" ]



