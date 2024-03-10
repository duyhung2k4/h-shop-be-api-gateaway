FROM golang:1.20.5

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 18880

CMD ["./main"]