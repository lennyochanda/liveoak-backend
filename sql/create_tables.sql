CREATE DATABASE IF NOT EXISTS users;

USE users;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id        CHAR(36) NOT NULL,
	username  VARCHAR(50) NOT NULL,
	email     VARCHAR(50) NOT NULL,
	password  VARCHAR(128) NOT NULL,
	createdAt DATETIME NOT NULL,
	updatedAt DATETIME NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users
    (id, username, email, password, createdAt, updatedAt)
VALUES
    ("2c9e4d53-c019-4655-9531-ea6eeeb9673d", "Lenny Ochanda", "lennyonduto384@gmail.com", "$SECRET", "2023-05-05 03:33:54", "2023-05-05 03:33:54"),
    ("b056c403-347d-4f87-b181-bd20fd6de987", "Wilson Ochanda", "wilsonochanda2000@gmail.com", "$SECRET", "2023-05-05 03:33:54", "2023-05-05 03:33:54");