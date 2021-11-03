CREATE TABLE IF NOT EXISTS users
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id    varchar(50)        NOT NULL UNIQUE,
    email      varchar(100)       NOT NULL UNIQUE,
    name       varchar(100)       NOT NULL DEFAULT '',
    picture    varchar(300)       NOT NULL,
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);

CREATE TABLE posts
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title      varchar(250)       NOT NULL DEFAULT '',
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now(),
    user_id    int                NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS tags
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name       varchar(50)        NOT NULL DEFAULT '',
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS post_tags
(
    post_id int NOT NULL,
    tag_id  int NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id),
    PRIMARY KEY (post_id, tag_id)
);
