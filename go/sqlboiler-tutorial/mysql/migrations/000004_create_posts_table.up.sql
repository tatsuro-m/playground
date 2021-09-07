CREATE TABLE posts
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title      varchar(250)       NOT NULL DEFAULT '',
    user_id    int                NOT NULL,
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now(),
    INDEX user_id_index (user_id),
    FOREIGN KEY fk_user_id (user_id) REFERENCES users (id)
);
