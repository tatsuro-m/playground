ALTER TABLE users ADD COLUMN created_at timestamp NOT NULL DEFAULT now();
ALTER TABLE users ADD COLUMN updated_at timestamp NOT NULL DEFAULT now();
