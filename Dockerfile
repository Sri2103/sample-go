# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Runner
FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8500

CMD ["./myapp"]