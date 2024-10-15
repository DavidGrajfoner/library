# Use the official Golang image to build the Go application
FROM golang:1.23.2

# Set the working directory in the container
WORKDIR /library

# Copy the current directory contents into the container at /app
COPY . .

# Download Go module dependencies
RUN go mod download

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go application
CMD ["go", "run", "main.go"]
