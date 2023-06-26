package routes

import (
	"net/http"

	"github.com/akifkadioglu/golang-websocket/controller/auth"
	"github.com/akifkadioglu/golang-websocket/controller/home"
	"github.com/akifkadioglu/golang-websocket/controller/socket"
	"github.com/akifkadioglu/golang-websocket/utils"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func CreateServer() *chi.Mux {
	oauth2GoogleConfig := InitConfigs()
	SetErrorMessages()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	/* socket için */
	go socket.Run()

	/* google için config'i ayarlıyorum */
	stateConfig := gologin.DebugOnlyCookieConfig

	/* herkese açık routelar */
	r.Group(func(r chi.Router) {
		r.Get("/", home.Home)
		r.Mount("/google/login", google.StateHandler(stateConfig, google.LoginHandler(oauth2GoogleConfig, nil)))
		r.Mount("/google/callback", google.StateHandler(stateConfig, google.CallbackHandler(oauth2GoogleConfig, http.HandlerFunc(auth.Login), nil)))
		r.HandleFunc("/socket/{name}", socket.SocketHandler)
	})

	/* yalnızca giriş yapmış kullanıcılar için oluşturulmuş route'lar */
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.JWTAuth()))
		r.Use(jwtauth.Authenticator)
		r.Get("/check", func(w http.ResponseWriter, r *http.Request) {})
	})

	return r
}
