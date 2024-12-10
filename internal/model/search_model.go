package model

type Search struct {
	Tracks   `json:"tracks"`
	Artists  `json:"artists"`
	Albums   `json:"albums"`
	Playlist `json:"playlist"`
	Shows    `json:"shows"`
}

type Tracks struct {
	Href     string  `json:"href"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}

type Artists struct {
	Href     string  `json:"href"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}
type Albums struct {
	Href     string  `json:"href"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}

type Playlist struct {
	Href     string  `json:"href"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}

type Shows struct {
	Href     string  `json:"href"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}

type Items struct {
	Id         *string `json:"id"`
	Name       *string `json:"name"`
	Popularity *int8   `json:"popularity"`
	Type       *string `json:"type"`

	Album struct {
		Id          *string `json:"id"`
		AlbumType   *string `json:"album_type"`
		TotalTracks *int8   `json:"total_tracks"`
	}
	Artist struct {
		Id   *string `json:"id"`
		Name *string `json:"name"`
		Type *string `json:"type"`
	}
}
