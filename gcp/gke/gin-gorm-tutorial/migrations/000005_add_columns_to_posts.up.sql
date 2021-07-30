ALTER TABLE posts ADD COLUMN created_at timestamp NOT NULL DEFAULT now();
ALTER TABLE posts ADD COLUMN updated_at timestamp NOT NULL DEFAULT now();
