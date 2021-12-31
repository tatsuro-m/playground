ALTER TABLE postal_codes
    ADD UNIQUE pcode_uk_uniq (code, prefecture_id, municipality_id, town_area_id);
