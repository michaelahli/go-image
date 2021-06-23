package usecases

import (
	"errors"
	"net/http"
	"strings"

	"github.com/michaelahli/Simple-Image-Transfer/server/src/models"
)

func (u *uc) SaveImage(w http.ResponseWriter, r *http.Request, auth *models.AuthModel) (string, int, error) {
	var metadata models.ImageMetadataModel

	// limit file upload to 8 megabyte
	r.Body = http.MaxBytesReader(w, r.Body, 8*1024*1024)

	// read file data
	image, headers, err := r.FormFile("file")
	if err != nil {
		return "Image is unprocessable", http.StatusUnprocessableEntity, err
	}
	defer image.Close()

	// detect image type
	fileHeader := make([]byte, 512)
	if _, err := image.Read(fileHeader); err != nil {
		return "Image type is undetectable", http.StatusUnprocessableEntity, err
	}

	// non image file is forbidden
	content_type := http.DetectContentType(fileHeader)
	if !strings.Contains(content_type, "image") {
		return "File type is unsupported. Please use image file instead.", http.StatusUnprocessableEntity, errors.New("Error Unsupported File Type")
	}

	metadata.Name = headers.Filename
	metadata.ContentType = content_type
	metadata.Size = headers.Size

	// insert image metadata to db
	err = u.repo.InsertOne(metadata, "image")
	if err != nil {
		return "Failed to save image metadata", http.StatusInsufficientStorage, nil
	}

	return "Successfully save image to database", http.StatusOK, nil
}
