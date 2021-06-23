package usecases

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/src/models"
	"gopkg.in/mgo.v2/bson"
)

func (u *uc) ProcessRequest(r *http.Request) (*models.AuthModel, string, int, error) {
	var (
		message string
		auth    models.AuthModel
		verify  models.AuthModel
	)

	authorization := r.Header.Get("Authorization")

	decoded, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		message = "Failed to decode authorization. Make sure to use base64 encoding."
		return nil, message, http.StatusUnprocessableEntity, err
	}

	err = json.Unmarshal([]byte(decoded), &auth)
	if err != nil {
		message = "Failed to unmarshal authorization"
		return nil, message, http.StatusInternalServerError, err
	}

	where := bson.M{"username": auth.Username}
	err = u.repo.FindOne(&verify, where, "user", nil)
	if err != nil {
		message = "User Not Found"
		return nil, message, http.StatusForbidden, err
	}

	return &auth, message, http.StatusOK, nil
}
