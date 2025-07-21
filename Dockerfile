FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

RUN adduser -D -s /bin/sh appuser

WORKDIR /root/

COPY --from=builder /app/main .

RUN chown appuser:appuser main

USER appuser

EXPOSE 3600

CMD ["./main"]