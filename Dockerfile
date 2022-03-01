FROM golang:latest

WORKDIR /app

COPY server.go ./

EXPOSE 8080

CMD ["go", "run", "./server.go"]