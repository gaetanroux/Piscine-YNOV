SELECT "FirstName", "LastName", "Company", "Total" InvoiceTotalPrice
FROM customers
INNER JOIN invoices ON customers."CustomerId" = invoices."CustomerId" WHERE "Company" IS NOT NULL