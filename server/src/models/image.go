package models

type ImageMetadataModel struct {
	Name        string `json:"name" bson:"name"`
	ContentType string `json:"content_type" bson:"content_type"`
	Size        int64  `json:"size" bson:"size"`
}
