package clients

import (
	"go.mongodb.org/mongo-driver/mongo"
	"phrase-api/models"
)

type IClient interface {
	Get(paging *models.Paging, sort *models.Sort) *mongo.Cursor
}
