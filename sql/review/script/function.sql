# case 文
SELECT Name,
       CASE
           WHEN CountryCode = 'JPN' THEN '日本の場所です'
           WHEN CountryCode <> 'JPN' THEN '日本以外の場所です'
           ELSE 'その他'
           END AS 'japan?'
FROM city;

SELECT Name,
       CASE
           WHEN CountryCode = 'JPN' THEN '日本の場所です'
           WHEN CountryCode <> 'JPN' THEN '日本以外の場所です'
           ELSE 'その他'
           END AS 'japan?'
FROM city
WHERE CountryCode = 'JPN';

# length
SELECT Name, LENGTH(Name) AS 'name length'
FROM city;

# trim
INSERT INTO city(ID, Name, CountryCode)
VALUES (999999, ' trim sample', 'JPN');

SELECT Name, TRIM(Name) AS 'trimmed name'
FROM city
WHERE Name = ' trim sample';

# replace
INSERT INTO city(ID, Name, CountryCode)
VALUES (99999, 'replace sample: replace here!', 'JPN');

UPDATE city
SET Name = REPLACE(Name, 'replace here!', 'replaced!')
WHERE Name = 'replace sample: replace here!';

# substring
# 文字列の一部を抽出して演算を行う
SELECT *
FROM city
WHERE SUBSTRING(Name, 1, 4) = 'toky';

SELECT *
FROM city
WHERE SUBSTRING(Name, 1, 4) LIKE '%ky%';

# round
SELECT CountryCode, Language, Percentage, ROUND(Percentage, -1) AS '小数第1位で四捨五入した結果'
FROM countrylanguage;

# trunc
SELECT CountryCode, Language, Percentage, TRUNCATE(Percentage, 0) AS '小数点以下を切り捨てた結果'
FROM countrylanguage;

# power（べき乗の計算）
SELECT Name, Population, POWER(Population, 2) AS '人口を2乗した数'
FROM city;

# current date,time,timestamp
# date や time 型のカラムが無かったので変だけど以下で代用
INSERT INTO city(CountryCode, Name)
VALUES ('JPN', CURRENT_DATE);

INSERT INTO city(CountryCode, Name)
VALUES ('JPN', CURRENT_TIME);

INSERT INTO city(CountryCode, Name)
VALUES ('JPN', CURRENT_TIMESTAMP);

# cast
SELECT Name, CONCAT(CAST(Population AS CHAR), '人') AS '人口'
FROM city;
