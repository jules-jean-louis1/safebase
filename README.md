# SafeBase

## Description du projet

SafeBase est un projet réalisé en Go, conçu pour permettre la gestion des backups et restaurations de bases de données PostgreSQL et MySQL. Le projet fournit une API REST qui permet aux utilisateurs de sauvegarder (dump) et restaurer des bases de données, soit à partir d'un conteneur Docker contenant la base de données, soit en téléchargeant directement les fichiers de backup depuis le frontend.

### Fonctionnalités clés :
- **Sauvegarde des bases de données** : Crée des dumps complets des bases de données.
- **Restauration des bases de données** : Permet la restauration à partir de fichiers de sauvegarde existants.
- **Suivi des opérations** : Garde un historique des backups et des restaurations effectués, avec des statuts de réussite ou d'échec.
- **Gestion des jobs planifiés (cron jobs)** : Automatisation des backups à intervalles réguliers.
- **Interface utilisateur** : Permet de télécharger et restaurer les backups directement via le frontend.

L'application s'appuie sur les outils `pg_dump` pour PostgreSQL et `mysqldump` pour MySQL afin d'assurer la sécurité et l'intégrité des sauvegardes.

## Initialisation du projet

Pour initialiser le projet, utilisez la commande suivante à la racine du projet pour créer tous les conteneurs nécessaires :
```bash
docker compose build
```

Démarrer les conteneurs :
```bash
docker compose up
```

Arrêter les conteneurs :
```bash
docker compose stop
```

(Re)build un seul conteneur :
```bash
docker compose build <container_name>
```