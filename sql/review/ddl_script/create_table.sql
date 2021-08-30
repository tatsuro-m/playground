CREATE TABLE mydb.users
(
    id   int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name varchar(50)                    NOT NULL DEFAULT ''
);

CREATE TABLE mydb.posts
(
    id      int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    title   varchar(100)                   NOT NULL,
    user_id int                            NOT NULL,
    index user_id_index (user_id),
    foreign key fkey_user_id (user_id) REFERENCES mydb.users (id)
);
