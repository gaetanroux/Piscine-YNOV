SELECT "FirstName", "LastName", strftime('%Y', "HireDate") - strftime('%Y', "BirthDate") as ApproximateAge
FROM employees
ORDER BY "BirthDate" DESC