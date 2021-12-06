SELECT
    *,
    COUNT(*) AS Nb_Rows
FROM
    (
        SELECT
            "Customer"."Country"
        FROM
            "Customer"
        UNION ALL
        SELECT
            "Employee"."Country"
        FROM
            "Employee"
        UNION ALL
        SELECT
            "Order"."ShipCountry"
        FROM
            "Order"
    )
GROUP BY
    "Country"