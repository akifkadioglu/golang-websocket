package routes

import (
	"net/http"

	"github.com/akifkadioglu/golang-websocket/controller/auth"
	"github.com/akifkadioglu/golang-websocket/controller/home"
	"github.com/akifkadioglu/golang-websocket/controller/socket"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer() *chi.Mux {
	oauth2GoogleConfig := InitConfigs()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	go socket.Run()
	r.Get("/", home.Home)

	stateConfig := gologin.DebugOnlyCookieConfig
	r.Mount("/google/login", google.StateHandler(stateConfig, google.LoginHandler(oauth2GoogleConfig, nil)))
	r.Mount("/google/callback", google.StateHandler(stateConfig, google.CallbackHandler(oauth2GoogleConfig, http.HandlerFunc(auth.Login), nil)))

	r.HandleFunc("/socket/{name}", socket.SocketHandler)
	return r
}
