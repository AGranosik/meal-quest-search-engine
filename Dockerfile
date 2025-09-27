FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies (only if go.mod or go.sum change)
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary for Linux (static binary)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ----------------------------
# STEP 2: Create a minimal runtime image
# ----------------------------
FROM alpine:3.20

# Add a non-root user for better security
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /app/main .

# Change ownership
RUN chown -R appuser /app

# Run as non-root user
USER appuser

# Expose the application port (optional, depends on your app)
EXPOSE 8080

# Run the Go binary
CMD ["./main"]