FROM golang:latest

RUN mkdir -p /go/src/simple-web-app
WORKDIR /go/src/simple-web-app
COPY ./gin /go/src/simple-web-app

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD go run hello.go
