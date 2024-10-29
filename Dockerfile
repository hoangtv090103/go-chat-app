# Step 1: Use the official Golang image to build the Go app
FROM golang:1.22.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .  
# Ensure this line is correctly copying all source files

# Build the Go app
RUN go build -o /app/go-chat-app

# Step 2: Use a smaller image to run the app
FROM alpine:latest
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/go-chat-app .

# Expose the port your app runs on
EXPOSE 3000

# Command to run the app
CMD ["/app/go-chat-app"]
