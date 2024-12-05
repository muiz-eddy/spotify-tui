package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"spotify-tui/internal"
)

func Search(query string, albumType string) ([]byte, error) {
	base, err := url.Parse(internal.BaseUrl)
	if err != nil {
		return nil, err
	}

	base.Path += "search"
	params := url.Values{}
	params.Add("q", query)
	params.Add("type", albumType)
	params.Add("market", "ES")
	base.RawQuery = params.Encode()

	endpoint := base.String()

	ctx := context.Background()
	client := CreateSpotifyClient(ctx)

	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch search results: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch search results: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return body, nil
}
