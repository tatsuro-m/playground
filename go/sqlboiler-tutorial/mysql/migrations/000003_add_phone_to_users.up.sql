ALTER TABLE users
    ADD COLUMN phone varchar(15) DEFAULT '' NOT NULL UNIQUE;
