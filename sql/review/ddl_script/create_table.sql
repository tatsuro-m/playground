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

# 外部キー制約が正しく設定されているかのテスト
INSERT INTO mydb.users(name)
VALUES ('first user');

INSERT INTO mydb.posts(title, user_id)
VALUES ('not exists user!', 9999);

INSERT INTO mydb.posts(title, user_id)
VALUES ('valid user!', 1);

SELECT *
FROM mydb.users
LEFT OUTER JOIN mydb.posts p ON users.id = p.user_id;
