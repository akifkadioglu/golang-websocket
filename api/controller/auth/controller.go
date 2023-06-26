package auth

import (
	"net/http"
	"time"

	"github.com/akifkadioglu/golang-websocket/database"
	"github.com/akifkadioglu/golang-websocket/ent"
	"github.com/akifkadioglu/golang-websocket/ent/user"
	"github.com/akifkadioglu/golang-websocket/models"
	"github.com/akifkadioglu/golang-websocket/utils"

	"github.com/dghubble/gologin/v2/google"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var jwtclaims models.JwtModel
	db := database.DBManager()

	/* kullanıcı bilgilerini google'dan alıyorum */
	userFromGoogle, err := google.UserFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/* kullanıcıyı arıyorum, yoksa yeni bi tane oluşturuyorum */
	var userClient *ent.User
	userClient, err = db.User.Query().
		Where(user.Email(userFromGoogle.Email)).
		First(r.Context())

	if ent.IsNotFound(err) {

		userClient, err = db.User.
			Create().
			SetEmail(userFromGoogle.Email).
			SetPicture(userFromGoogle.Picture).
			SetName(userFromGoogle.Name).
			Save(r.Context())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	/* JWT'yi ayarlıyorum */
	jwtclaims.ID = userClient.ID.String()
	jwtclaims.Email = userClient.Email
	jwtclaims.Name = userClient.Name
	jwtclaims.Picture = *userClient.Picture
	jwtclaims.Time = time.Now().String()

	/* token oluşturuyorum */
	mapData, _ := utils.StructToMap(jwtclaims)
	_, tokenString, _ := utils.JWTAuth().Encode(mapData)

	/* tokenla giriş yapıyorum */
	w.Header().Add("Authorization", "Bearer "+tokenString)
	http.Redirect(w, r, "https://talk-secretly.netlify.app/token/"+tokenString, http.StatusMovedPermanently)
}
