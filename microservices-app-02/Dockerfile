FROM golang:1.20 AS builder

WORKDIR /app

# Copy the source code into the container
COPY . .

# Install dependencies
RUN go mod download

# Build the binary
RUN go build -o /app/server ./cmd/server

FROM alpine:3.14

# Copy the binary into the container
COPY --from=builder /app/server /usr/local/bin/server

# Set the binary as the entrypoint
ENTRYPOINT ["/usr/local/bin/server"]
# docker build -t microservices-app-b:1.0.0 --build-arg GO_VERSION=1.20 .
