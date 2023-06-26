package user

import (
	"net/http"
	"time"

	"github.com/akifkadioglu/golang-websocket/database"
	"github.com/akifkadioglu/golang-websocket/models"
	"github.com/akifkadioglu/golang-websocket/utils"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func ChangeName(w http.ResponseWriter, r *http.Request) {
	db := database.DBManager()
	var jwtclaims models.JwtModel
	var input ChangeNameBody

	err := render.Decode(r, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := utils.GetUser(r)
	uuid, err := uuid.Parse(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.User.UpdateOneID(uuid).
		SetName(input.Name).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwtclaims.ID = user.ID
	jwtclaims.Email = user.Email
	jwtclaims.Name = input.Name
	jwtclaims.Picture = user.Picture
	jwtclaims.Time = time.Now().String()

	mapData, _ := utils.StructToMap(jwtclaims)
	_, tokenString, _ := utils.JWTAuth().Encode(mapData)

	render.JSON(w, r, map[string]string{
		"message": "The name changed to " + input.Name,
		"token":   tokenString,
	})
}
