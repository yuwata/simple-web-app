FROM golang:latest

RUN mkdir -p /go/src/users
WORKDIR /go/src/users
COPY src/users /go/src/users

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD go run main.go
