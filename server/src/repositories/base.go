package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	mongodb *mongo.Database
}

type Repository interface {
	FindOne(res, where interface{}, tableName string, opts *options.FindOneOptions) error
	FindAll(where interface{}, tableName string) (*mongo.Cursor, error)
	InsertOne(data interface{}, tableName string) error
}

func SetupRepository(mongodb *mongo.Database) Repository {
	return &repo{mongodb: mongodb}
}
