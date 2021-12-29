CREATE TABLE IF NOT EXISTS prefectures
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name       varchar(50)        NOT NULL UNIQUE,
    name_roma  varchar(50)        NOT NULL UNIQUE,
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
