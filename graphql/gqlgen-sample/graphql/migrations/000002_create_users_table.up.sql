CREATE TABLE IF NOT EXISTS users
(
    id         int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    #"AgU2cRWNdrcDbE4zDbeKBc9mfAr1"
    user_id    varchar(50)        NOT NULL UNIQUE,
    email      varchar(100)       NOT NULL UNIQUE,
    name       varchar(100)       NOT NULL DEFAULT '',
    picture    varchar(300)       NOT NULL,
    created_at DATETIME           NOT NULL DEFAULT now(),
    updated_at DATETIME           NOT NULL DEFAULT now()
);
