ALTER TABLE postal_codes
    DROP CONSTRAINT code,
    ADD UNIQUE pcode_uk_uniq (code, prefecture_id, municipality_id, town_area_id);
