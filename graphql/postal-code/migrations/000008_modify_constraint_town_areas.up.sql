ALTER TABLE town_areas
    DROP CONSTRAINT name,
    ADD UNIQUE uk_t_n_nr_m(name, name_roma, municipality_id);
