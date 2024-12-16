package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"spotify-tui/internal"
	"spotify-tui/internal/services"
	"testing"
)

// TODO: to understand more about this
func TestSearch(t *testing.T) {

	mockResponseData := `{
		"tracks": {
			"href": "mock_href",
			"total": 1,
			"items": [
				{
					"id": "track1",
					"name": "Mock Track",
					"popularity": 80,
					"type": "track",
					"album": {
						"id": "album1",
						"album_type": "single",
						"total_tracks": 1
					},
					"artist": {
						"id": "artist1",
						"name": "Mock Artist",
						"type": "artist"
					}
				}
			]
		}
	}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search" {
			t.Fatalf("Expected URL /search, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, mockResponseData)
	}))

	defer mockServer.Close()

	internal.BaseUrl = mockServer.URL
	ctx := context.Background()
	mockClient := &http.Client{}

	query := "taylor swift"
	searchType := "track"

	result, err := services.Search(ctx, query, searchType, mockClient)

	if err != nil {
		t.Fatalf("Expected no error, go: %v", err)
	}

	expected := "Mock Track"

	if len(result.Tracks.Items) == 0 || *result.Tracks.Items[0].Name != expected {
		t.Errorf("Expected track name '%s', got '%s'", expected, *result.Tracks.Items[0].Name)
	}

}
