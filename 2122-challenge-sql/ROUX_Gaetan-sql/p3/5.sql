SELECT "ShipName" AS 'Clients', "ShipVia", "ShipCountry" AS 'Country'
FROM "Order"
WHERE "ShipCountry" = 'Brazil'
AND "ShipVia" = 3
GROUP BY "CustomerId"
ORDER BY "ShipVia" DESC