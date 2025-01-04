# Use the official Go image as the base
FROM golang:1.23-alpine AS builder

LABEL version="1.0.0"
LABEL description="API Rest Service provides an unique global ID"
LABEL license="MIT"
LABEL source="https://github.com/chrix75/id-generator"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o main ./api

# Use a scratch image for a smaller final image size
FROM scratch

# Copy the built application binary
COPY --from=builder /app/main /app/main

# Expose the port the application listens on
EXPOSE 8080

ENV GIN_MODE=release

# Set the entrypoint to run the application
ENTRYPOINT ["/app/main"]