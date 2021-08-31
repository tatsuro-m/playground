ALTER TABLE users
    ADD COLUMN email      varchar(30) DEFAULT '' NOT NULL UNIQUE,
    ADD COLUMN created_at timestamptz            NOT NULL DEFAULT current_timestamp,
    ADD COLUMN updated_at timestamptz            NOT NULL DEFAULT current_timestamp;

