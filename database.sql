-- Active: 1753800467741@@127.0.0.1@3306@golang_database
CREATE TABLE customer
(
    id      VARCHAR(100) NOT NULL,
    name    VARCHAR(100) NOT NULL
    PRIMARY KEY(id)
)ENGINE = InnoDb;

DELETE FROM customer;

ALTER Table customer
    ADD COLUMN email    VARCHAR(100),
    ADD COLUMN balance  INTEGER DEFAULT 0,
    ADD COLUMN rating   DOUBLE  DEFAULT 0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    add COLUMN birth_date DATE,
    add COLUMN married  BOOLEAN DEFAULT false;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES ('eko', 'Eko', 'eko@gmail.com', 100000, 5.0, '1999-9-9', true);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES ('budi', 'Budi', 'budi@gmail.com', 100000, 5.0, '1999-9-9', true);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES ('joko', 'Joko', NULL, 100000, 5.0, NULL, true);

CREATE TABLE user
(
username VARCHAR(100) NOT NULL,
password VARCHAR(100) NOT NULL,
PRIMARY KEY (username)
) ENGINE = InnoDB;

INSERT INTO user(username, password) VALUES('admin', 'admin')

CREATE TABLE comments
(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT,
    PRIMARY KEY (id)
) ENGINE InnoDB;



