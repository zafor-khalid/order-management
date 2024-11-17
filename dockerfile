# Use an official Go image with version 1.23
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker's layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Navigate to the directory containing the main file
WORKDIR /app/cmd/api

# Build the Go application
RUN go build -o /app/main .

# Set the entry point to the compiled binary
CMD ["/app/main"]
