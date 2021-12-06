SELECT UPPER(LastName) AS LastNameUpper, "FirstName"
FROM customers
GROUP BY "LastName"
having LENGTH(LastName) > 5 AND LENGTH(FirstName) > 5 
ORDER BY "LastName" ASC;