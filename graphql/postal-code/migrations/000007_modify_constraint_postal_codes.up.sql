ALTER TABLE postal_codes
    DROP CONSTRAINT number,
    ADD UNIQUE pcode_uk_uniq (number, prefecture_id, municipality_id, town_area_id);
