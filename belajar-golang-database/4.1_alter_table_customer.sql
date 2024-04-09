DELETE FROM customer;

ALTER TABLE customer
    ADD COLUMN email        VARCHAR(255),
    ADD COLUMN balance      INTEGER     DEFAULT 0,
    ADD COLUMN rating       DOUBLE      DEFAULT 0.0,
    ADD COLUMN created_at   TIMESTAMP   DEFAULT CURRENT_TIMESTAMP(),
    ADD COLUMN birth_date   DATE,
    ADD COLUMN married      BOOLEAN     DEFAULT false;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES ('eko', 'Eko', 'eko@gmail.com', 1000000, 90.0, '1990-10-10', true),
    ('wendy', 'Wendy', 'wendy@gmail.com', 500000, 100.0, '2000-10-10', true);