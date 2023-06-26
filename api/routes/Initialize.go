package routes

import (
	"errors"
	"os"

	"github.com/go-chi/jwtauth/v5"
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
		RedirectURL:  "https://socket-nwnt.onrender.com/google/callback",
		Endpoint:     googleOAuth2.Endpoint,
		Scopes:       []string{"profile", "email"},
	}
}

func SetErrorMessages() {
	jwtauth.ErrUnauthorized = errors.New("404 page not found")
	jwtauth.ErrExpired = errors.New("404 page not found")
	jwtauth.ErrNBFInvalid = errors.New("404 page not found")
	jwtauth.ErrIATInvalid = errors.New("404 page not found")
	jwtauth.ErrNoTokenFound = errors.New("404 page not found")
	jwtauth.ErrAlgoInvalid = errors.New("404 page not found")
}
