package controllers

import (
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/src/usecases"
)

type ctrl struct {
	uc usecases.Usecase
}

type Controller interface {
	// @location : image.go
	ImageHandler(w http.ResponseWriter, r *http.Request)

	// @location : check.go
	CheckMetadata(w http.ResponseWriter, r *http.Request)
}

func SetupController(uc usecases.Usecase) Controller {
	return &ctrl{uc: uc}
}
