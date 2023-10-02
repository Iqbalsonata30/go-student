FROM golang:1.21-alpine

WORKDIR  /app

COPY . .

RUN go test -v ./test

RUN go build -o ./bin/go-student

EXPOSE 3000

CMD ./bin/go-student
