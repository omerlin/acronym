# Stage 1: Build the Go binary
FROM docker.io/library/golang:1.20-alpine AS builder

# Install git for go mod download
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary for linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Run the binary in a minimal image
FROM scratch

# Install ca-certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set the working directory in the second stage
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application will run on
EXPOSE 3000

# Command to run the application
CMD ["./main", "-http"]
