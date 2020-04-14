package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"phrase-api/clients"
	"phrase-api/models"
)

type PhraseRepository struct {
	client clients.IClient
}

func NewPhraseRepository(client clients.IClient) *PhraseRepository {
	return &PhraseRepository{client: client}
}

func (self *PhraseRepository) GetPhraseStatsList(paging *models.Paging, sort *models.Sort) []models.PhraseStats {
	var result []models.PhraseStats
	var cursor *mongo.Cursor = self.client.Get(paging, sort)

	ctx := context.TODO()
	for cursor.Next(ctx) {
		var phrase models.PhraseStats
		err := cursor.Decode(&phrase)
		if err != nil {
			log.Fatal(err)
		} else {
			result = append(result, phrase)
		}
	}

	defer cursor.Close(ctx)

	return result
}
