-- Création de la base de données
CREATE DATABASE IF NOT EXISTS bibliotheque_en_ligne;
USE bibliotheque_en_ligne;

-- Création de la table auteurs
CREATE TABLE auteurs (
    id_auteur INT AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    nationalite VARCHAR(50)
);

-- Création de la table categories
CREATE TABLE categories (
    id_categorie INT AUTO_INCREMENT PRIMARY KEY,
    nom_categorie VARCHAR(50) NOT NULL
);

-- Création de la table livres
CREATE TABLE livres (
    id_livre INT AUTO_INCREMENT PRIMARY KEY,
    titre VARCHAR(150) NOT NULL,
    annee_publication INT,
    id_auteur INT,
    id_categorie INT,
    FOREIGN KEY (id_auteur) REFERENCES auteurs(id_auteur),
    FOREIGN KEY (id_categorie) REFERENCES categories(id_categorie)
);

-- Insertion d'exemples de données dans la table auteurs
INSERT INTO auteurs (nom, nationalite) VALUES
('J.K. Rowling', 'Royaume-Uni'),
('George Orwell', 'Royaume-Uni'),
('Gabriel Garcia Marquez', 'Colombie');

-- Insertion d'exemples de données dans la table categories
INSERT INTO categories (nom_categorie) VALUES
('Fiction'),
('Science-fiction'),
('Classiques');

-- Insertion d'exemples de données dans la table livres
INSERT INTO livres (titre, annee_publication, id_auteur, id_categorie) VALUES
('Harry Potter à l''école des sorciers', 1997, 1, 1),
('1984', 1949, 2, 2),
('Cent ans de solitude', 1967, 3, 3);