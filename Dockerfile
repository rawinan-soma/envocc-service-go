# Multi-stage build
FROM golang:1.24-alpine AS builder

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the main API server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

# Build the migrator/seeder
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/reset-seed/reset-seed ./cmd/reset-seed/reset-seed.go

# Final stage
FROM alpine:latest

# Install ca-certificates and curl for healthcheck
RUN apk --no-cache add ca-certificates curl tzdata

ENV TZ=Asia/Bangkok

# Create app directory
WORKDIR /root/

# Copy binaries from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/cmd/reset-seed/reset-seed ./cmd/reset-seed/
COPY config.yaml .
# COPY --from=builder /app/config ./config/
# COPY --from=builder /app/assets ./assets/

# Make binaries executable
RUN chmod +x main
RUN chmod +x cmd/reset-seed/reset-seed

# Expose port
EXPOSE 3600

# Default command (can be overridden)
CMD ["./main"]