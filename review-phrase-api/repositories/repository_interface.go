package repositories

import "phrase-api/models"

type IRepository interface {
	GetPhraseStatsList(paging models.Paging, sort models.Sort) []models.PhraseStats
}
