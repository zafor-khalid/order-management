# Use the latest official Golang image to build the application
FROM golang:1.23.2-alpine3.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o zafor-dev main.go

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/zafor-dev .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./zafor-dev"]
