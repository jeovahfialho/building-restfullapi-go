# Use the official Go image as the base image
FROM golang:1.17-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and its dependencies
COPY go.mod .
COPY go.sum .

# Download all dependencies
RUN go mod download

# Copy the source code into the working directory
COPY . .

# Build the application to a binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/

# Stage 2: Execution
FROM scratch

# Copy the binary and the script to the root directory of the scratch image
COPY --from=builder /app/main /main
COPY --from=builder /app/wait-for-it.sh /wait-for-it.sh

# Optionally, set the working directory if you need to execute the binary in a specific directory
# WORKDIR /app

# Ensure wait-for-it.sh has the correct permissions before building the image
# Command to execute the binary
CMD ["/main"]
