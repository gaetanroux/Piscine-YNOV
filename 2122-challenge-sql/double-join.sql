SELECT DISTINCT tracks."Name" TrackName, playlists."Name" PlaylistName
FROM playlist_track
INNER JOIN playlists ON playlists."PlaylistId" = playlist_track."PlaylistId"
INNER JOIN tracks ON playlist_track."TrackId" = tracks."TrackId"
WHERE playlists."Name" = 'TV Shows' 
LIMIT 100