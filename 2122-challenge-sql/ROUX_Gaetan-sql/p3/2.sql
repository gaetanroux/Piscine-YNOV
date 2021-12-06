SELECT
    "LastName",
    "FirstName",
    "TerritoryId",
    "Territory"."TerritoryDescription"
FROM
    "EmployeeTerritory"
    INNER JOIN "Territory" ON "EmployeeTerritory"."TerritoryId" = "Territory"."Id"
    INNER JOIN "Employee" ON "EmployeeTerritory"."EmployeeId" = "Employee"."Id"
WHERE
    "Employee"."Id" = 4;