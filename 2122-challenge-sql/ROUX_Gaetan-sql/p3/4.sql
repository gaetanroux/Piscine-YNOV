SELECT
    "ShipName" AS 'Clients',
    "OrderDate" AS 'Date',
    "ShipCity" AS 'City'
FROM
    "Order"
WHERE
    "City" = 'London'
    AND "ShippedDate" <= 2014