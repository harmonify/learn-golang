CREATE TABLE IF NOT EXISTS user (
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (username)
) ENGINE = InnoDB;

INSERT INTO
    user(username, password)
VALUES
    ('admin', 'admin');