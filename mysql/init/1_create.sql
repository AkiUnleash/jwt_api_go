CREATE DATABASE golang_db;
use golang_db;
CREATE TABLE accounts (
  id INT(11) AUTO_INCREMENT NOT NULL,
  name VARCHAR(64) NOT NULL,
  email CHAR(30) NOT NULL,
  password CHAR(30) NOT NULL,
  PRIMARY KEY (id)
);
INSERT INTO
  accounts (name, email, password)
VALUES
  ("Jupiter", "aaa@example.com", "aaa");
INSERT INTO
  accounts (name, email, password)
VALUES
  ("Charlotte", "bbb@example.com", "bbb");
INSERT INTO
  accounts (name, email, password)
VALUES
  ("Wing", "ccc@example.com", "ccc");