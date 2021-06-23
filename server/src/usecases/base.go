package usecases

import (
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/src/models"
	"github.com/michaelahli/Simple-Image-Transfer/server/src/repositories"
)

type uc struct {
	repo repositories.Repository
}

type Usecase interface {
	ProcessRequest(r *http.Request) (*models.AuthModel, string, int, error)
	SaveImage(w http.ResponseWriter, r *http.Request, auth *models.AuthModel) (string, int, error)
	FetchAllMetadata() ([]models.ImageMetadataModel, string, int, error)
}

func SetupUsecase(repo repositories.Repository) Usecase {
	return &uc{repo: repo}
}
