CREATE TABLE IF NOT EXISTS postal_codes
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    number     int                NOT NULL UNIQUE,
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
