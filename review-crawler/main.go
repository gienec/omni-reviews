package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/antchfx/htmlquery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gopkg.in/jdkato/prose.v2"
	"log"
	"strings"
	"time"
)

func main() {
	processNextpage := true
	page := 1
	for processNextpage == true {
		log.Println(fmt.Sprintf("=========== STARTED processing page %v ===========", page))

		url := fmt.Sprint("https://apps.shopify.com/omnisend/reviews?page=", page)
		processNextpage = processPageReviews(url)

		log.Println(fmt.Sprintf("=========== FINISHED processing page %v ===========", page))

		page = page + 1
	}
}

func processPageReviews(url string) bool {
	reviews := loadReviews(url)
	client := getConnection()
	collection := client.Database("omni").Collection("phrases")

	for i, review := range reviews {
		if isReviewProcessed(review) {
			continue
		}

		doc, err := prose.NewDocument(review)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Processing review", i, ":", review)

		for _, sentence := range doc.Sentences() {
			log.Println("Sentence:", sentence.Text)

			phrases := toPhrases(sentence.Text, 3)

			for j, phrase := range phrases {
				phrase = strings.ToLower(phrase)
				log.Println("Phrase ", j, ":", phrase)

				opts := options.Update().SetUpsert(true)
				filter := bson.D{{"phrase", phrase}}
				update := bson.D{{"$inc", bson.D{{"frequency", 1}}}}
				_, err := collection.UpdateOne(context.Background(), filter, update, opts)

				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	defer client.Disconnect(context.TODO())

 	return len(reviews) > 0
}

func isReviewProcessed(review string) bool {
	ctx := context.TODO()

	client := getConnection()
	collection := client.Database("omni").Collection("reviews")

	hash := getHash(review)
	filter := bson.D{{"hash", hash}}
	cursor, _ := collection.Find(ctx, filter)

	var status bool
	if cursor.Next(ctx) {
		status = true
	} else {
		status = false
		collection.InsertOne(ctx, bson.D{{"hash", hash}})
	}

	cursor.Close(ctx)
	defer client.Disconnect(ctx)

	log.Println(fmt.Sprintf("Review %s processed: %v", hash, status))

	return status
}

func getConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.TODO(), 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://omni:<password>@omnicomments-t1xlh.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func loadReviews(url string) []string {
	doc, _ := htmlquery.LoadURL(url)
	reviews := htmlquery.Find(doc, "//*[contains(concat(\" \", normalize-space(@class), \" \"), \" review-listing \")]")

	var result []string
	for _, review := range reviews {
		paragraph := htmlquery.Find(review, "//p")
		result = append(result, paragraph[0].FirstChild.Data)
	}

	return result
}

func toPhrases(sentence string, phraseWordCount int) [] string {
	var phrases []string
	words := strings.Split(sentence, " ")
	sentenceWordCount := len(words)

	for i, _ := range words {
		var phraseWords []string

		for nextWordIndex := i; nextWordIndex < sentenceWordCount; nextWordIndex++ {
			var word = words[nextWordIndex]
			if word != "" {
				phraseWords = append(phraseWords, word)
			}

			if len(phraseWords) == phraseWordCount {
				phrases = append(phrases, strings.Join(phraseWords, " "))
				break
			}
		}
	}

	return phrases
}

var hash = sha1.New()
func getHash(text string) string {
	hash.Reset()
	hash.Write([]byte(text))
	byteArrayHash := hash.Sum(nil)

	return fmt.Sprintf("%x", byteArrayHash)
}