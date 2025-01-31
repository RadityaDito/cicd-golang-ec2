# Stage 1: Build the application
FROM golang:1.22 as builder

WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy all project files
COPY . .

# Build the Go binary with static linking for Alpine compatibility
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

WORKDIR /app

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Ensure the binary is executable
RUN chmod +x /app/main

# Expose application port
EXPOSE 8080

# Run the application and load environment variables
CMD ["sh", "-c", "export $(cat .env | xargs) && /app/main"]
