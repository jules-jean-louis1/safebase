# Utiliser l'image officielle Node.js comme étape de build
FROM node:18 AS build

# Définir le répertoire de travail
WORKDIR /app

COPY . .

RUN npm install


RUN npm run build -- --configuration build

# Utiliser une image Nginx pour servir l'application Angular
FROM nginx:alpine

COPY --from=build /app/dist/ /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exposer le port 80 pour Nginx
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]
