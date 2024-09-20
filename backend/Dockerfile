# Use Go image
FROM golang:1.23.0

# Install necessary packages including MySQL client
RUN apt-get update && \
    apt-get install -y default-mysql-client postgresql-client && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Setup application
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Expose port and define entrypoint
EXPOSE 8080
CMD ["air"]
