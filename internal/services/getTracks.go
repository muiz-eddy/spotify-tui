package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"spotify-tui/internal"
	"spotify-tui/internal/model"
)

/* TODO:
* Add tests for Get Track
 */
func GetSingleTrack(ctx context.Context, id string, client *http.Client) (model.Tracks, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if client == nil {
		client = CreateSpotifyClient(ctx)
	}

	base, _ := url.Parse(internal.BASEURL)

	base.Path += "tracks"
	params := url.Values{
		"id": {id},
	}
	base.RawQuery = params.Encode()

	endpoint := base.String()

	resp, err := client.Get(endpoint)

	if err != nil {
		return model.Tracks{}, fmt.Errorf("failed to fetch search results: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return model.Tracks{}, fmt.Errorf("unexpected Status Code %d", resp.StatusCode)
	}

	var result model.Tracks

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return model.Tracks{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// TODO: Add endpoint of Get Several Tracks
func GetMultipleTracks() {

}
