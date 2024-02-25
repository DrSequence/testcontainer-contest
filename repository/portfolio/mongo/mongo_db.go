package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb(uri, database, collection string) (*mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: "root",
		Password: "example",
	}))
	if err != nil {
		return nil, err
	}
	db := client.Database(database)
	col := db.Collection(collection)
	return col, err
}
