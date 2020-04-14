package repositories

import (
	"phrase-api/clients"
	"phrase-api/models"
)

type PhraseRepository struct {
	client clients.IClient
}

func NewPhraseRepository(client clients.IClient) *PhraseRepository {
	return &PhraseRepository{client: client}
}

func (self *PhraseRepository) GetPhraseStatsList(paging models.Paging, sort models.Sort) []models.PhraseStats {

	var result []models.PhraseStats
	var s struct{}
	items := self.client.Get(10, s)


	return result
}
