ALTER TABLE postal_codes
    MODIFY COLUMN number varchar(7) NOT NULL UNIQUE;
