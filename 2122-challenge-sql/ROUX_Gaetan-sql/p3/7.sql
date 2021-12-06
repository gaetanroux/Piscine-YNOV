SELECT
    "OrderId",
    SUM(
        "UnitPrice" * "Quantity" * (100 - "Discount") / 100
    ) TotalPrice
FROM
    "OrderDetail"
GROUP BY
    "OrderId"
HAVING
    SUM(
        "UnitPrice" * "Quantity" * (100 - "Discount") / 100
    ) < (
        SELECT
            AVG(
                "UnitPrice" * "Quantity" * (100 - "Discount") / 100
            ) AveragePrice
        FROM
            "OrderDetail"
    )