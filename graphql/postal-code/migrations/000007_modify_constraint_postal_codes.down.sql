ALTER TABLE postal_codes
    ADD UNIQUE (number),
    DROP CONSTRAINT pcode_uk_uniq;
