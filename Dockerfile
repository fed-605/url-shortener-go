FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build  -o /app/exe ./cmd/url-shortener/main.go

CMD ["./app/exe"]
