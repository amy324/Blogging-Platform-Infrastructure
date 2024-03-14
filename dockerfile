FROM golang:1.22.1-alpine3.19

WORKDIR /app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire project directory into the Docker image context
COPY . .

# Build the executable binary and place it in the /usr/local/bin/ directory
RUN go build -v -o /usr/local/bin/app .

# Print the contents of /usr/local/bin and current working directory when the container starts
CMD ["sh", "-c", "ls -l /usr/local/bin && pwd && ./app"]
