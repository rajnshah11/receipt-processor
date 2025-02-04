# Stage 1: Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Install necessary tools and dependencies
RUN apk add --no-cache git

# Copy Go modules and source code into container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o receipt-processor .

# Stage 2: Runtime stage
FROM alpine:latest

WORKDIR /app

# Add a non-root user for security
RUN adduser -D appuser
USER appuser

# Copy executable from build stage
COPY --from=builder /app/receipt-processor .

# Expose port for the application
EXPOSE 8000

# Command to run the executable
CMD ["./receipt-processor"]
