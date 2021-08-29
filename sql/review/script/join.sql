SELECT *
FROM city
         JOIN country ON country.Code = city.CountryCode;

SELECT *
FROM city
         JOIN country c ON c.Code = city.CountryCode;

# inner join
SELECT *
FROM country
        INNER JOIN city c on country.Code = c.CountryCode
WHERE Code = 'JPN';

# outer join
SELECT *
FROM country
         LEFT OUTER JOIN city c ON country.Code = c.CountryCode
WHERE Code = 'JPN';

SELECT country.Name, Country.Population, c.Name AS 'city の名称'
FROM country
         LEFT OUTER JOIN city c ON country.Code = c.CountryCode
WHERE Code = 'JPN';

SELECT *
FROM country
         RIGHT OUTER JOIN city c ON country.Code = c.CountryCode
WHERE Code = 'JPN';
