# Step 1: Build
FROM golang:1.23.4 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o myapp

# Step 2: Run
FROM debian:bullseye-slim
WORKDIR /root/
COPY --from=builder /app/myapp .
EXPOSE 8080
CMD ["./myapp"]
