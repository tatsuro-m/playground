# 和集合（２つの結果の足し合わせ）を取得する UNION 関数なので、本来であれば同じような形のテーブルに利用する
SELECT ID, Name, CountryCode FROM city
UNION
SELECT ID, Name, CountryCode FROM city;

SELECT CountryCode, Language FROM countrylanguage
UNION
SELECT CountryCode, Name FROM city;
