CREATE TABLE IF NOT EXISTS post_tags
(
    post_id int NOT NULL,
    tag_id  int NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id),
    PRIMARY KEY (post_id, tag_id)
)
