# Stage 1: Build
FROM golang:1.22 AS builder

WORKDIR /app

# Copy source code
COPY . .

# Build the Go binary
RUN go mod init hello && go build -o server main.go

# Stage 2: Run
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
