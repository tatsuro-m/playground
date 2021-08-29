# sum
SELECT SUM(Population) AS '人口の合計'
FROM city
WHERE CountryCode = 'JPN';

# count
# null も含める
SELECT COUNT(*) AS '全レコード数'
FROM city;

# 指定したカラムが null なレコードは含めない
SELECT COUNT(Name) AS 'name が not null なレコード数'
FROM city;

# group by
SELECT Name, SUM(Population) AS '地域ごとの人口合計'
FROM city
GROUP BY Name;
