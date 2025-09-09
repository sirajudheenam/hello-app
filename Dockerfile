# 1️⃣ Build stage
FROM golang:1.25.1 AS builder

# Install git for go get
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
# RUN go build -o hello-server main.go
# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o hello-server main.go

# 2️⃣ Final stage: minimal Alpine image
FROM alpine:3.20

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/hello-server .

# Create a logs directory
RUN mkdir -p /app/logs

# Set environment variable for log file
ENV LOG_FILE=/app/logs/server.json.log

# Expose port
EXPOSE 8080

# Run the server
CMD ["./hello-server"]
# CMD ["sh", "-c", "./hello-server >> $LOG_FILE 2>&1"]