SELECT
    "OrderId",
    COUNT("Id") NbDifferentProducts,
    SUM(
        "UnitPrice" * "Quantity" * (100 - "Discount") / 100
    ) TotalPrice
FROM
    "OrderDetail"
WHERE
    "OrderId" IN (
        SELECT
            "OrderId"
        FROM
            "OrderDetail"
        GROUP BY
            "OrderId"
        HAVING
            COUNT("OrderId") > 5
    )
GROUP BY
    "OrderId";