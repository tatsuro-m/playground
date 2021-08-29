SELECT *
FROM city
WHERE Name = 'Rafah' AND CountryCode = 'JPN';

SELECT *
FROM city
WHERE Name = 'Rafah' AND CountryCode = 'PSE';

SELECT *
FROM city
WHERE CountryCode <> 'JPN';

SELECT *
FROM city
WHERE CountryCode <> 'JPN' OR CountryCode = 'JPN';

SELECT *
FROM city
WHERE CountryCode = 'JPN' OR CountryCode = 'USA';
