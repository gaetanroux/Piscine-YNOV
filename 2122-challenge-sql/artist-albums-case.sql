SELECT DISTINCT artists."Name", COUNT(*) AS NbAlbums,
CASE
     WHEN COUNT(*) = 1 THEN 'Unproductive'
     WHEN COUNT(*) > 1 AND COUNT(*) < 10 THEN 'Productive'
     ELSE 'Very Productive'
END
'IsProductive'
FROM artists
INNER JOIN albums ON artists."ArtistId" = albums."ArtistId"
GROUP BY artists."ArtistId"
LIMIT 100
