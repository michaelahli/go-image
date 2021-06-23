package controllers

import (
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/config/utils"
)

func (c *ctrl) CheckMetadata(w http.ResponseWriter, r *http.Request) {
	_, msg, code, err := c.uc.ProcessRequest(r)
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	res, msg, code, err := c.uc.FetchAllMetadata()
	if err != nil {
		utils.BasicResponse(w, msg, code, nil)
		return
	}

	utils.BasicResponse(w, msg, code, res)
	return
}
