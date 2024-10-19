# Build stage
FROM golang:alpine AS builder

# Install necessary packages
RUN apk add --no-cache git

# Set the working directory
WORKDIR /go/src/app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /main /main

# Expose the application port
EXPOSE 50000

# Command to run the application
CMD ["/main"]