FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first so Docker can reuse this layer when code changes.
COPY go.mod go.sum ./
RUN go mod download

# Build the service as a small Linux binary for the final runtime image.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

WORKDIR /app

# Ship only the compiled binary to keep the runtime image minimal.
COPY --from=builder /app/app .

EXPOSE 8080

# Start the HTTP service when the container launches.
CMD ["./app"]
