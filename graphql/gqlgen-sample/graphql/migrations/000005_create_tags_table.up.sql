CREATE TABLE IF NOT EXISTS tags
(
    id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name varchar(50)      NOT NULL DEFAULT '',
    created_at DATETIME   NOT NULL DEFAULT now(),
    updated_at DATETIME   NOT NULL DEFAULT now()
)
