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

func Search(ctx context.Context, query string, searchType string, client *http.Client) (model.Search, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if client == nil {
		client = CreateSpotifyClient(ctx)
	}

	base, err := url.Parse(internal.BASEURL)
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
