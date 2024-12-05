package services

import (
	"context"
	"spotify-tui/internal"
)

func Search() {
	q := "taylor swift"
	endpoint := internal.BaseUrl + "search" + "?=" + q

	ctx := context.Background()
	client := CreateSpotifyClient(ctx)

	resp, err := client.Get(endpoint)
}
