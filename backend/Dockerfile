FROM golang:1.23.0 AS builder

# Installer les dépendances nécessaires pour la construction
RUN apt-get update && \
    apt-get install -y wget gnupg && \
    echo "deb http://apt.postgresql.org/pub/repos/apt bookworm-pgdg main" > /etc/apt/sources.list.d/pgdg.list && \
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    apt-get update && \
    apt-get install -y postgresql-client-16 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Configurer l'application
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Étape 2 : Exécution
FROM alpine:latest

# Installer les dépendances nécessaires pour l'exécution
RUN apk --no-cache add ca-certificates

# Copier le binaire depuis l'étape de construction
COPY --from=builder /app/main /app/main

# Définir le répertoire de travail
WORKDIR /app

# Exposer le port nécessaire
EXPOSE 8080

# Définir le point d'entrée
CMD ["./main"]
