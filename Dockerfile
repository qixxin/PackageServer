FROM golang:latest
ADD ./server.go /go/src/app

RUN mkdir /app
WORKDIR /app

RUN go build -o server app/server

CMD ["/app/server"]
EXPOSE 8080

