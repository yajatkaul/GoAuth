package inits

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func SetupDatabase() {
	var err error
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))

	DB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")


	// Making fields unqiue
	collection := DB.Database("testdb").Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1}, 
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("Could not create unique index:", err)
	}
	log.Println("Unique index created for username field.")
}