package model

type Search struct {
}

type Tracks struct {
}

type Album struct {
}

type Playlist struct{}

type Shows struct{}

type Items struct {
	Id         string `json:"id"`
	Popularity int8   `json:"popularity"`
	Type       string `json:"type"`

	Album struct {
		Id          string `json:"id"`
		AlbumType   string `json:"album_type"`
		TotalTracks int8   `json:"total_tracks"`
	}
	Artist struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	}
}
