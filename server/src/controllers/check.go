package controllers

import (
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/config/utils"
)

/*
	@function : Fetch all data in database
*/

func (c *ctrl) CheckMetadata(w http.ResponseWriter, r *http.Request) {
	// Authenticate User
	_, msg, code, err := c.uc.ProcessRequest(r)
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	// Fetch Data
	res, msg, code, err := c.uc.FetchAllMetadata()
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	utils.BasicResponse(w, msg, code, res)
	return
}
