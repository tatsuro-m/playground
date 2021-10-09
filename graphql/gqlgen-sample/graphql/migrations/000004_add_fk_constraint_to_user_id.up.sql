ALTER TABLE posts
    ADD FOREIGN KEY fk_user_id(user_id) REFERENCES users (id);
