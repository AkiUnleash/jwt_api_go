CREATE DATABASE golang_db;
use golang_db;

CREATE TABLE users (
    id INT(11) AUTO_INCREMENT NOT NULL,
    name VARCHAR(64) NOT NULL,
    email CHAR(30) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (name, email) VALUES ("Jupiter", "aaa@example.com");
INSERT INTO users (name, email) VALUES ("Charlotte", "bbb@example.com");
INSERT INTO users (name, email) VALUES ("Wing", "ccc@example.com");