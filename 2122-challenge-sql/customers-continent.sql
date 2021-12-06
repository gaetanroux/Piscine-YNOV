SELECT "FirstName", "LastName", "Country",
  CASE 
    WHEN "Country" IN ('Brazil','Canada','USA','Argentina','Chile') THEN 'America'
    WHEN "Country"= 'India' THEN 'Asia'
    WHEN "Country"= 'Australia' THEN 'Oceania'
    ELSE 'Europe'
   END
	'Continent'
FROM customers