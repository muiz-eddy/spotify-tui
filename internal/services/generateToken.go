package services

import (
	"fmt"
	"net/http"
	"net/url"
)

func GenerateToken() (string, error) {
	formData := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {""},
		"client_secret": {""},
	}
	resp, err := http.PostForm("https://accounts.spotify.com/api/token", formData)

	if err != nil {
		return "", fmt.Errorf("failed to send request %v", err)
	}
	defer resp.Body.Close()
}
