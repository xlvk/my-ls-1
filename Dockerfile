# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Install the 'file' command
RUN apt-get update && apt-get install -y file && rm -rf /var/lib/apt/lists/*

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o my-ls-1 main/main.go

# Set execute permission for the binary
RUN chmod +x my-ls-1

# Override the default command to start a shell
CMD ["/bin/bash"]

