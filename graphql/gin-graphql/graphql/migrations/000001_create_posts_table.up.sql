CREATE TABLE posts
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title      varchar(250)       NOT NULL DEFAULT '',
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
