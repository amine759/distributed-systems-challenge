# Use the official Golang image as the base image
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files and dependencies
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
