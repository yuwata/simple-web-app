FROM golang:latest

RUN mkdir -p /go/src/hello
WORKDIR /go/src/hello
COPY src/hello /go/src/hello

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD go run hello.go
