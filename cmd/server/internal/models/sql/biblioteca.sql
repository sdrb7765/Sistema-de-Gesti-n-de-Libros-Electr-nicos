CREATE DATABASE biblioteca_virtual;
USE biblioteca_virtual;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100)
);

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100),
    author VARCHAR(100)
);

INSERT INTO books (title, author) VALUES
('El Quijote', 'Miguel de Cervantes'),
('Cien Años de Soledad', 'Gabriel García Márquez'),
('La Odisea', 'Homero');
