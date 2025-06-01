# Stage 1: Build
FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o order-service .

# Stage 2: Run
FROM alpine:latest

WORKDIR /root

# Copy built binary and .env file from the build stage
COPY --from=build /app/order-service /order-service
COPY --from=build /app/.env .env

COPY --from=build /app/serviceaccountgcp.json /app/credentials.json

# Set environment variable to tell the app where to find credentials
# Make sure to mount credentials at runtime if needed
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json

# Expose port 8080 for the application to listen
EXPOSE 8080

# Set the default command to run the application
CMD ["/order-service"]
