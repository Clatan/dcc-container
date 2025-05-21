# Stage 1: Build Go binary
FROM golang:1.20-alpine AS builder

# Install git (diperlukan untuk mengunduh Gin)
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files dan unduh dependensi terlebih dahulu (agar caching efisien)
COPY go.mod go.sum ./
RUN go mod download

# Salin semua source code dan build
COPY . .
RUN go build -o app .

# Stage 2: Runtime image
FROM alpine:latest

# Copy file binary dari builder
COPY --from=builder /app/app /app

# Expose port yang digunakan Gin (default 8080)
EXPOSE 8080

# Jalankan aplikasi
CMD ["/app"]
