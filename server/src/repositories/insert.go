package repositories

import "context"

func (r *repo) InsertOne(data interface{}, tableName string) error {
	_, err := r.mongodb.Collection(tableName).InsertOne(context.TODO(), data, nil)
	if err != nil {
		return err
	}
	return nil
}
