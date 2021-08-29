# % で任意の0文字以上の文字列を表す
SELECT *
FROM city
WHERE Name LIKE '%koha%';

SELECT *
FROM city
WHERE Name LIKE 'koha%';

#  _ で任意の1文字を表す
SELECT *
FROM city
WHERE Name LIKE '%koham__';
