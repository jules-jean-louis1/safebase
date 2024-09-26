# Utiliser l'image officielle Node.js comme étape de build
FROM node:18 AS build

# Définir le répertoire de travail
WORKDIR /usr/src/app

# Copier package.json et package-lock.json pour installer les dépendances
COPY package*.json ./

# Installer les dépendances
RUN npm install

# Copier tout le code source
COPY . .

# Construire l'application Angular pour la production
RUN npm run build -- --configuration production

# Utiliser une image Nginx pour servir l'application Angular
FROM nginx:alpine

# Copier les fichiers construits de l'étape précédente dans le répertoire Nginx
COPY --from=build /usr/src/app/dist/ /usr/share/nginx/html

# Copier le fichier de configuration Nginx personnalisé
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exposer le port 80 pour Nginx
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]
