ALTER TABLE users
    ADD COLUMN email      varchar(30) DEFAULT '' NOT NULL UNIQUE,
    ADD COLUMN created_at DATETIME               NOT NULL DEFAULT now(),
    ADD COLUMN updated_at DATETIME               NOT NULL DEFAULT now();

