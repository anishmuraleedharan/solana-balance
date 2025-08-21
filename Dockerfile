FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o solana-api
EXPOSE 8080
CMD ["./solana-api"]