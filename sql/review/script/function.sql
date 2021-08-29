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
VALUES (999999,' trim sample', 'JPN');

SELECT Name, TRIM(Name) AS 'trimmed name'
FROM city
WHERE Name = ' trim sample';

# replace
INSERT INTO city(ID, Name, CountryCode)
VALUES (99999,'replace sample: replace here!', 'JPN');

UPDATE city
SET Name = REPLACE(Name, 'replace here!', 'replaced!')
WHERE Name = 'replace sample: replace here!';

