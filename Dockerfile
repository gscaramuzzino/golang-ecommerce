# ---- Build Stage ----
FROM golang:1.19-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux 

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -a -installsuffix cgo -o app ./cmd/ecommerce-api

# ---- Run Stage ----
FROM alpine:latest

WORKDIR /app

# Copy the Pre-built binary file from the previous stage. 
COPY --from=builder /build/app .

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run when starting the container
CMD ["./app"]
