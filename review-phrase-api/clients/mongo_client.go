package clients

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	phraseCollection *mongo.Collection
}

func NewMongoClient(phraseCollection *mongo.Collection) *MongoClient {
	return &MongoClient{
		phraseCollection,
	}
}

func (self *MongoClient) Get(limit int, filter interface{}) []interface{} {
	cursor, _ := self.phraseCollection.Find(context.Background())

	var result []interface{}

	return result
}