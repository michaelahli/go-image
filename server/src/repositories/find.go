package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) FindAll(where interface{}, tableName string) (*mongo.Cursor, error) {
	result, err := r.mongodb.Collection(tableName).Find(context.TODO(), where, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) FindOne(res, where interface{}, tableName string, opts *options.FindOneOptions) error {
	result := r.mongodb.Collection(tableName).FindOne(context.TODO(), where)
	if result.Err() != nil {
		return result.Err()
	}
	result.Decode(res)
	return nil
}
