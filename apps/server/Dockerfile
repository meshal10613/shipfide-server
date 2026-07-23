# Build stage
FROM golang:1.26.3-alpine AS builder

# Install certificates and git
RUN apk update && apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /app

# Copy dependency manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server cmd/api/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Expose server port
EXPOSE 5000

# Start server
ENTRYPOINT ["./server"]
