package services

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

type HTTPService struct{}

func CreateSpotifyClient(ctx context.Context) *http.Client {
	bearerToken, _ := GenerateToken()
	token := &oauth2.Token{
		AccessToken: bearerToken,
	}
	tokenSource := oauth2.StaticTokenSource(token)

	return oauth2.NewClient(ctx, tokenSource)
}
