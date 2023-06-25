package routes

import (
	"os"

	"golang.org/x/oauth2"
	googleOAuth2 "golang.org/x/oauth2/google"
)

type Config struct {
	ClientID     string
	ClientSecret string
}

func InitConfigs() *oauth2.Config {
	googleConfig := &Config{
		ClientID:     os.Getenv("ClientID"),
		ClientSecret: os.Getenv("ClientSecret"),
	}
	return &oauth2.Config{
		ClientID:     googleConfig.ClientID,
		ClientSecret: googleConfig.ClientSecret,
		RedirectURL:  "172.168.10.243:3000/google/callback",
		Endpoint:     googleOAuth2.Endpoint,
		Scopes:       []string{"profile", "email"},
	}
}
