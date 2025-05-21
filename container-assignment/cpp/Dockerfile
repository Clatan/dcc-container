# Stage 1: Build di Alpine (gunakan static linking)
FROM alpine:latest AS builder

RUN apk add --no-cache g++ musl-dev

WORKDIR /app
COPY . .

# Build dengan static linking
RUN g++ -static -o hello main.cpp

# Stage 2: Runtime Alpine
FROM scratch

COPY --from=builder /app/hello /hello

CMD ["/hello"]