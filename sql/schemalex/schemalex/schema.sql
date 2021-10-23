CREATE TABLE hoge
(
    id    INTEGER     NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (id),
    title varchar(50) NOT NULL DEFAULT ''
);

CREATE TABLE posts
(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title varchar(50) NOT NULL DEFAULT ''
);
