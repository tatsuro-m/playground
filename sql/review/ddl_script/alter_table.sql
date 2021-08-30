ALTER TABLE mydb.users
    ADD COLUMN email varchar(20) NOT NULL UNIQUE,
    ADD COLUMN phone varchar(15) NOT NULL UNIQUE;

ALTER TABLE mydb.users
    DROP email,
    DROP phone;
