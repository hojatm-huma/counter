FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main

EXPOSE 8080

ENV DB_HOST="host.docker.internal" DB_USER="ccp" DB_PASSWORD="pass123" DB_NAME="counter"

CMD ["./main"]


