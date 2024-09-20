-- Création de la table artistes dans la base de données par défaut
CREATE TABLE artistes (
    id_artiste SERIAL PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    pays_origine VARCHAR(50)
);

-- Création de la table genres dans la base de données par défaut
CREATE TABLE genres (
    id_genre SERIAL PRIMARY KEY,
    nom_genre VARCHAR(50) NOT NULL
);

-- Création de la table albums dans la base de données par défaut
CREATE TABLE albums (
    id_album SERIAL PRIMARY KEY,
    titre VARCHAR(100) NOT NULL,
    annee_sortie INT,
    id_artiste INT REFERENCES artistes(id_artiste),
    id_genre INT REFERENCES genres(id_genre)
);

-- Insertion d'exemples de données dans la table artistes
INSERT INTO artistes (nom, pays_origine) VALUES 
('The Beatles', 'Royaume-Uni'), 
('Miles Davis', 'États-Unis'), 
('Daft Punk', 'France');

-- Insertion d'exemples de données dans la table genres
INSERT INTO genres (nom_genre) VALUES 
('Rock'), 
('Jazz'), 
('Électronique');

-- Insertion d'exemples de données dans la table albums
INSERT INTO albums (titre, annee_sortie, id_artiste, id_genre) VALUES 
('Abbey Road', 1969, 1, 1), 
('Kind of Blue', 1959, 2, 2), 
('Discovery', 2001, 3, 3);