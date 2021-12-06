SELECT customers."FirstName"||' '|| UPPER(customers."LastName") AS FullName, SUM(invoices."Total") as AllInvoices 
FROM customers
INNER JOIN invoices ON customers."CustomerId" = invoices."CustomerId"
GROUP BY customers."CustomerId"
HAVING SUM(invoices."Total") > 38
ORDER BY customers."FirstName" ASC
