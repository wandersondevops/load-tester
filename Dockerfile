
# Use the official Golang image to create a build artifact.
FROM golang:1.18 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the initialization script
COPY init_go_mod.sh ./

# Make the script executable
RUN chmod +x init_go_mod.sh

# Run the initialization script to create go.mod and go.sum
RUN ./init_go_mod.sh

# Copy the source code to the container
COPY . .

# Build the Go app with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o load-tester

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Install bash to help with debugging if needed
RUN apk add --no-cache bash

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/load-tester .

# Verify the binary exists and is executable
RUN ls -l load-tester

# Command to run the executable
ENTRYPOINT ["./load-tester"]
    