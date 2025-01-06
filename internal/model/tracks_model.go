package model

type Tracks struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	TrackNumber   int    `json:"track_number"`
	Type          string `json:"type"`
	TrackDuration string `json:"duration_ms"`
	IsPlayable    string `json:"is_playable"`
	TrackAlbum    `json:"album"`
	TrackArtist   []TrackArtist `json:"artist"`
}

type TrackAlbum struct {
	AlbumType   string `json:"album_type"`
	TotalTracks string `json:"total_tracks"`
	ReleaseDate string `json:"release_date"`
}

type TrackArtist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
