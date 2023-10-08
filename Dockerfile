FROM golang:latest

# Set working directory in container
WORKDIR /app

COPY . .

# Build
RUN go build -o main .

EXPOSE 80

CMD ["./main"]
