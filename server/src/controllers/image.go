package controllers

import (
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/config/utils"
)

func (c *ctrl) ImageHandler(w http.ResponseWriter, r *http.Request) {
	auth, msg, code, err := c.uc.ProcessRequest(r)
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	msg, code, err = c.uc.SaveImage(w, r, auth)
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	utils.BasicResponse(w, msg, code, nil)
	return
}
