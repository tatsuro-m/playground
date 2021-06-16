ALTER TABLE posts ADD COLUMN user_id integer REFERENCES users(id) NOT NULL;
