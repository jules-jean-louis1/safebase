# Étape de compilation
FROM golang:1.23.1 AS build

# On configure le répertoire de travail
WORKDIR /app

# On copie les fichiers de dépendances go mod et go sum pour profiter du cache docker
COPY go.mod go.sum ./

# Installation des dépendances
RUN go mod download

# On copie le reste du projet
COPY . .

# On construit l'exécutable Go avec des paramètres pour un binaire statique et optimisé pour Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /safebase-backend

# Étape finale : utilisation d'une image légère
FROM alpine:latest

# On ajoute le certificat CA pour éviter les erreurs SSL lors des appels réseau
RUN apk --no-cache add ca-certificates

# On définit l'utilisateur à non-root pour des raisons de sécurité (optionnel)
USER 1000:1000

# On copie l'exécutable construit dans l'étape précédente
COPY --from=build /safebase-backend /safebase-backend

# On expose le port 8080
EXPOSE 8080

# Commande de lancement de l'application
ENTRYPOINT ["/safebase-backend"]