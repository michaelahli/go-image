package usecases

import (
	"context"
	"net/http"

	"github.com/michaelahli/Simple-Image-Transfer/server/src/models"
	"gopkg.in/mgo.v2/bson"
)

func (u *uc) FetchAllMetadata() ([]models.ImageMetadataModel, string, int, error) {
	var metadatas []models.ImageMetadataModel

	// fetch all data in image collection
	where := bson.M{}
	res, err := u.repo.FindAll(where, "image")
	if err != nil {
		return nil, "Failed to fetch metadata", http.StatusInternalServerError, err
	}

	// copy all data to proper struct
	for res.Next(context.TODO()) {
		var metadata models.ImageMetadataModel
		err = res.Decode(&metadata)
		if err != nil {
			return nil, "Failed to decode metadata", http.StatusInternalServerError, err
		}
		// append all data into array
		metadatas = append(metadatas, metadata)
	}

	return metadatas, "success", http.StatusOK, nil
}
