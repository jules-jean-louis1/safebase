# Use Go image
FROM golang:1.23-alpine

# Install necessary packages including MariaDB client
RUN apk update && apk add --no-cache \
    docker-cli \
    git \
    postgresql-client \
    mariadb-client \
    libc6-compat

# Verify versions
RUN go version && mysqldump --version

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
