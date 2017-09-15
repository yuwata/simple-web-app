FROM golang:latest

RUN mkdir -p /go/src/restful
WORKDIR /go/src/restful
COPY src/restful /go/src/restful

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD go run main.go
