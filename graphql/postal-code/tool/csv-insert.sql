USE dev;

SET GLOBAL local_infile = on;
LOAD DATA LOCAL INFILE 'KEN_ALL_ROME.CSV'
    INTO TABLE prefectures
    CHARACTER SET sjis
    FIELDS TERMINATED BY ','
    (@1, @2, @3, @4, @5, @6)
    SET name = @2, name_roma = @6;
