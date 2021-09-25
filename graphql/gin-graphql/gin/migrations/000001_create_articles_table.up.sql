CREATE TABLE IF NOT EXISTS articles
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    author     varchar(50)        NOT NULL DEFAULT '',
    title      varchar(250)       NOT NULL DEFAULT '',
    content    varchar(400)       NOT NULL DEFAULT '',
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
