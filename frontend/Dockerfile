# Utiliser l'image officielle Node.js pour le développement
FROM node:18

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /usr/src/app

# Copier package.json et package-lock.json dans le conteneur
COPY package*.json ./

# Installer les dépendances Angular
RUN npm install

# Copier tout le code source dans le conteneur
COPY . .

# Exposer le port utilisé par le serveur de développement Angular
EXPOSE 4200

# Commande pour démarrer le serveur Angular en mode développement
CMD ["npm", "start"]
