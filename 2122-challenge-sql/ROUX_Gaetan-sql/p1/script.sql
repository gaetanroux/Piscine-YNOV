CREATE TABLE clients
(
    id_clients INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    nom VARCHAR(100),
    prenom VARCHAR(100),
    email VARCHAR(255),
    id_chambres INT NOT NULL,
    );
CREATE TABLE chambres
(
    id_chambres INT PRIMARY KEY NOT NULL,
    id_clients INT NOT NULL,
    nb_places INT NOT NULL,
    prix INT NOT NULL
    );
CREATE TABLE reservation
(
    id_clients INT PRIMARY KEY NOT NULL,
    id_chambres INT NOT NULL,
    tp_sejours INT NOT NULL,
    date_reservation DATE,
    prix INT NOT NULL
    );
CREATE TABLE restaurants
(
    date_menu DATE PRIMARY KEY NOT NULL,
    entree VARCHAR(255),
    plat VARCHAR(255),
    dessert VARCHAR(255),
    prix INT NOT NULL
    );