package controllers

import (
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/config/utils"
)

/*
	@function : accept image data and save it's metadata to db
*/
func (c *ctrl) ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Authenticate user
	auth, msg, code, err := c.uc.ProcessRequest(r)
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	// Save Image to DB
	msg, code, err = c.uc.SaveImage(w, r, auth)
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	utils.BasicResponse(w, msg, code, nil)
	return
}
