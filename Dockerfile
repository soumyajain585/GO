# Use an official Go runtime as a parent image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .
COPY go.mod .
COPY go.sum .

# Build the Go binary
RUN go build -o myapp

# Expose a port that the application will listen on
EXPOSE 8080

# Define the command to run the application
CMD ["./myapp"]
