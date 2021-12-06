SELECT artists."Name"
FROM albums
INNER JOIN artists ON albums."ArtistId" = artists."ArtistId"
GROUP by albums."ArtistId"
HAVING COUNT (albums."ArtistId") >= 4
ORDER BY artists."Name" DESC