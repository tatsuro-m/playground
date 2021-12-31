ALTER TABLE postal_codes
    ADD UNIQUE (code),
    DROP CONSTRAINT pcode_uk_uniq;
