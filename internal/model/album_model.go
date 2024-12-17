package model

type Album struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	ReleaseDate      string   `json:"release_date"`
	AlbumType        string   `json:"album_type"`
	TotalTracks      int      `json:"total_tracks"`
	AvailableMarkets []string `json:"available_markets"`
	Artists          []Artist `json:"artists"`
	Popularity       int      `json:"popularity"`
}

type Artist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
