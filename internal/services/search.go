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

// Search TODO: Add Testing for this endpoint
func Search(query string, searchType string) (model.Search, error) {
	base, err := url.Parse(internal.BaseUrl)
	if err != nil {
		return model.Search{}, err
	}

	base.Path += "search"
	params := url.Values{}
	params.Add("q", query)
	params.Add("type", searchType)
	params.Add("market", "ES")
	params.Add("limit", "10")
	base.RawQuery = params.Encode()

	endpoint := base.String()

	ctx := context.Background()
	client := CreateSpotifyClient(ctx)

	resp, err := client.Get(endpoint)
	if err != nil {
		return model.Search{}, fmt.Errorf("failed to fetch search results: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return model.Search{}, fmt.Errorf("failed to fetch search results: %d", resp.StatusCode)
	}

	var result model.Search

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return model.Search{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return result, nil
}
