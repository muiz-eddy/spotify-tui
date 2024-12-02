package services

import (
	"fmt"
	"net/http"
	"net/url"
	"spotify-tui/internal"
)

func GenerateToken() (string, error) {

	formData := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {internal.GetEnv("SPOTIFY_CLIENT_ID")},
		"client_secret": {internal.GetEnv("SPOTIFY_CLIENT_SECRET")},
	}
	resp, err := http.PostForm("https://accounts.spotify.com/api/token", formData)

	if err != nil {
		return "", fmt.Errorf("failed to send request %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected Status Code %d", resp.StatusCode)
	}

	var result map[string]interface{}
	accessToken, ok := result["access_token"].(string)

	if !ok {
		return "", fmt.Errorf("access token not found in response")
	}

	return accessToken, nil
}
