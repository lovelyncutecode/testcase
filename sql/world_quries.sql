SELECT country.Name,city.Population AS CapitalPopulation
FROM world.country JOIN world.city
ON country.Capital=city.ID
ORDER BY city.Population DESC 
LIMIT 5;

SELECT round(sum(Population*Percentage/100)) AS EnglishInEurPopulation 
FROM world.countrylanguage INNER JOIN world.country ON countrylanguage.CountryCode=country.`Code` 
WHERE Language='English' AND Continent='Europe';

SELECT of.`CountryCode` ,of.c AS OfficialLang,unof.c AS UnofficialLang
FROM (SELECT `CountryCode`, count(countrylanguage.IsOfficial) AS c 
       FROM  world.countrylanguage 
       WHERE countrylanguage.IsOfficial='T'
       GROUP BY `CountryCode`) AS of,
       (SELECT `CountryCode`, count(countrylanguage.IsOfficial) AS c 
       FROM world.countrylanguage 
       WHERE countrylanguage.IsOfficial='F'
       GROUP BY `CountryCode`) AS unof
WHERE of.`CountryCode`=unof.`CountryCode` AND unof.c>=(of.c*2) AND of.c>=2;  
