CREATE TABLE IF NOT EXISTS town_areas
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name       varchar(50)        NOT NULL,
    name_roma  varchar(50)        NOT NULL,
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
