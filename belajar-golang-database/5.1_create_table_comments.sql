CREATE TABLE IF NOT EXISTS comments (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    comment TEXT,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

INSERT INTO
    comments(email, comment)
VALUES
    ('admin@example.com', 'Hello all');
