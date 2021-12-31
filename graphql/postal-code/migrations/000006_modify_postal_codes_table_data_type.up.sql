ALTER TABLE postal_codes
    MODIFY COLUMN code varchar(7) NOT NULL UNIQUE;
