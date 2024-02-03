FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main cmd/password_generator/main.go

EXPOSE 3000

CMD ["./main"]
