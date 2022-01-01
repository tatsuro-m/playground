ALTER TABLE town_areas
    DROP CONSTRAINT uk_t_n_nr_m,
    ADD UNIQUE (name, municipality_id);
