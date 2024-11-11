package pkg

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	oauthConfig *oauth2.Config
)

func InitializeOAuth() {
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

func GetLoginURL(state string) string {
	return oauthConfig.AuthCodeURL(state)
}

func HandleOAuthCallback(r *http.Request) (*oauth2.Token, error) {
	code := r.URL.Query().Get("code")
	if code == "" {
		return nil, errors.New("code not found in callback")
	}

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging token: %v", err)
		return nil, err
	}

	return token, nil
}
