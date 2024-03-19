# Use an official Go runtime as the base image
FROM golang:1.21-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY main.go .

# Build the Go application
RUN go build -o /app/app main.go

# Create a new lightweight image for running the application
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app/

# Copy the binary from the build stage to the final stage
COPY --from=build --chown=nobody:root /app/app .

USER nobody

# Command to run the application
CMD ["/app/app"]
