FROM golang:alpine

# Installer les dépendances nécessaires, y compris Docker CLI, PostgreSQL client et MySQL client
RUN apk update && apk add --no-cache \
    docker-cli \
    git \
    postgresql-client \
    mysql-client 

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Installer Air pour le rechargement automatique pendant le développement
RUN go install github.com/air-verse/air@latest

# Copier go.mod et go.sum dans le conteneur
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source dans le conteneur
COPY . .

# Construire l'application Go
RUN go build -o main .

# Exposer le port 8080 pour l'accès extérieur
EXPOSE 8080

# Commande pour exécuter Air (outil de rechargement en développement)
CMD ["air"]
