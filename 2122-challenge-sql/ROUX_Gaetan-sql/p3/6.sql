SELECT
    "OrderId",
    SUM(
        "UnitPrice" * "Quantity" * (100 - "Discount") / 100
    ) TotalPrice
FROM
    "OrderDetail"
WHERE "OrderId" = 10260