package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	db         string
	collection string
}

func NewDB() MongoDb {
	return MongoDb{
		db:         "broker",
		collection: "messages",
	}
}

type Storage interface {
	insert(entity interface{}) error
	insertMany(entities []interface{}) error
}

func NewMongoDBClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("%s: %s", "mongodb new client error", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("%s: %s", "mongodb connect error", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("%s: %s", "mongodb ping error", err)
	}

	return client
}

func (m MongoDb) insert(entity interface{}) error {
	client := NewMongoDBClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Database(m.db).Collection(m.collection).InsertOne(ctx, entity)

	return err
}

func (m MongoDb) insertMany(entities []interface{}) error {
	client := NewMongoDBClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Database(m.db).Collection(m.collection).InsertMany(ctx, entities)

	return err
}
