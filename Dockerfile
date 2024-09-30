# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o /app/main ./cmd/server/main.go

# Expose port 8087 to the outside world
EXPOSE 8087

# Command to run the executable
CMD ["/app/main"]
