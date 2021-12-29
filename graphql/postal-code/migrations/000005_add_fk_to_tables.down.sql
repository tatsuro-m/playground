ALTER TABLE municipalities
    DROP FOREIGN KEY municipalities_ibfk_1,
    DROP COLUMN prefecture_id;

ALTER TABLE town_areas
    DROP FOREIGN KEY town_areas_ibfk_1,
    DROP COLUMN municipality_id;

ALTER TABLE postal_codes
    DROP FOREIGN KEY postal_codes_ibfk_1,
    DROP COLUMN prefecture_id,
    DROP FOREIGN KEY postal_codes_ibfk_2,
    DROP COLUMN municipality_id,
    DROP FOREIGN KEY postal_codes_ibfk_3,
    DROP COLUMN town_area_id;
