SELECT DISTINCT CountryCode
FROM city;

SELECT DISTINCT CountryCode
FROM city
WHERE Name IN ('tokyo', 'osaka');

SELECT DISTINCT CountryCode, Name
FROM city
WHERE Name IN ('tokyo', 'osaka');
