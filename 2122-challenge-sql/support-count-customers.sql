SELECT employees."FirstName"||' '|| UPPER(employees."LastName") AS FullName, COUNT(customers."SupportRepId") as NumberOfCustomers
FROM employees
INNER JOIN customers ON employees."EmployeeId" = customers."SupportRepId"
GROUP by employees."FirstName"
ORDER BY customers."SupportRepId" DESC