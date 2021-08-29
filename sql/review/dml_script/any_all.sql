# サブクエリを使ったパターン
# ANY 内の式でどれか１つでも一致するものがあれば真として扱われる
SELECT *
FROM city
WHERE Name = ANY (SELECT Name FROM city WHERE CountryCode = 'JPN');

SELECT *
FROM city
WHERE Name = ALL (SELECT Name FROM city WHERE CountryCode = 'JPN');

SELECT *
FROM city
WHERE Name = ALL (SELECT Name FROM city WHERE Name = 'tokyo');
