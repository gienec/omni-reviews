package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	. "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"phrase-api/clients"
	"phrase-api/handlers"
	"phrase-api/repositories"
)

func main() {
	router := gin.Default()
	api := router.Group("/v1")
	{
		api.GET("/phrases", getPhraseHandler().Handle)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}

func getPhraseHandler() handlers.IHandler {
	return handlers.PhraseHandler{
		Repository: getRepository(),
	}
}

func getRepository() repositories.IRepository {
	return repositories.NewPhraseRepository(getClient())
}

func getClient() clients.IClient {
	return clients.NewMongoClient(GetCollection(
		"mongodb+srv://omni:<password>@omnicomments-t1xlh.mongodb.net?retryWrites=true&w=majority",
		"omni",
		"phrases"))
}

func GetCollection(connectionString string , databaseName string, collectionName string) *mongo.Collection {
	ctx := context.Background()

	client, err := mongo.NewClient(
		Client().ApplyURI(connectionString),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	database := client.Database(databaseName)
	collection := database.Collection(collectionName)
	return collection
}