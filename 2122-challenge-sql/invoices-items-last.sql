SELECT tracks."Name"
FROM tracks
Where tracks."TrackId" = (SELECT invoice_items."TrackId"
FROM invoices
INNER JOIN invoice_items ON invoices."InvoiceId" = invoice_items."InvoiceId"
WHERE invoices."InvoiceId" = (SELECT invoices."InvoiceId"
FROM invoices
INNER JOIN invoice_items ON invoices."InvoiceId" = invoice_items."InvoiceId"
WHERE "InvoiceDate" = (
    SELECT invoices."InvoiceDate"
FROM invoices
ORDER BY "InvoiceDate" DESC
LIMIT 1
  )
  )
  )
