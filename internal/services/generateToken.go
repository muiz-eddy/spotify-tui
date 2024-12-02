package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"spotify-tui/internal"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

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
		return "", fmt.Errorf("unexpected Status Code %d", resp.StatusCode)
	}

	var result tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return result.AccessToken, nil
}
