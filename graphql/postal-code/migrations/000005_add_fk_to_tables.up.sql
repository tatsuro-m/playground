ALTER TABLE town_areas
    ADD COLUMN municipality_id int NOT NULL,
    ADD FOREIGN KEY (municipality_id) REFERENCES municipalities (id),
    ADD UNIQUE (name, municipality_id);


ALTER TABLE municipalities
    ADD COLUMN prefecture_id int NOT NULL,
    ADD UNIQUE (name, prefecture_id),
    ADD FOREIGN KEY (prefecture_id) REFERENCES prefectures (id);


ALTER TABLE postal_codes
    ADD COLUMN prefecture_id   int NOT NULL,
    ADD FOREIGN KEY (prefecture_id) REFERENCES prefectures (id),
    ADD COLUMN municipality_id int NOT NULL,
    ADD FOREIGN KEY (municipality_id) REFERENCES municipalities (id),
    ADD COLUMN town_area_id    int NOT NULL,
    ADD FOREIGN KEY (town_area_id) REFERENCES town_areas (id);
