SELECT *
FROM city
WHERE Name IN ('osaka', 'tokyo');

SELECT *
FROM city
WHERE Name NOT IN ('osaka', 'tokyo');

SELECT *
FROM city
WHERE Name NOT IN ('osaka', 'tokyo') AND CountryCode = 'JPN';
