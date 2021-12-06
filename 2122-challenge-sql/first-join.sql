SELECT "Title" AlbumName, "Name" ArtistName
FROM artists
INNER JOIN albums ON artists."ArtistId" = albums."ArtistId"
LIMIT 100