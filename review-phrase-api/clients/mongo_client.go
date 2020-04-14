package clients

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"phrase-api/models"
)

type MongoClient struct {
	phraseCollection *mongo.Collection
}

func NewMongoClient(phraseCollection *mongo.Collection) *MongoClient {
	return &MongoClient{
		phraseCollection,
	}
}

func (self *MongoClient) Get(paging *models.Paging, sort *models.Sort) *mongo.Cursor {
	options := options.Find()

	if sort != nil {
		var order int
		if order = 1; sort.Order == "desc" {
			order = -1
		}

		options.SetSort(bson.D{{sort.FieldName,order}})
	}

	options.SetSkip(paging.Offset)
	options.SetLimit(paging.Limit)

	cursor, _ := self.phraseCollection.Find(context.Background(), bson.D{}, options)
	return cursor
}