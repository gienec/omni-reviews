package models

type PhraseStats struct {
	Phrase string `bson:"phrase" json:"phrase"`
	Frequency int `bson:"frequency" json:"frequency"`
}