SELECT "InvoiceId", tracks."Name" AS InvoiceItem, invoice_items."UnitPrice" AS UnitPrice
FROM invoice_items
INNER JOIN tracks ON invoice_items."TrackId" = tracks."TrackId"
Where invoice_items."InvoiceId" = 10
ORDER BY tracks."Name"