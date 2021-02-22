package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"portservice/domain"
	"time"
)
func getMongoURI() string {
	return "mongodb://"+getEnv("MONGO_HOST", "localhost")+":27017"
}

func Save(ports []*domain.Port) error {
	log.Printf("saving %v", len(ports))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("port").Collection("ports")

	var ui []interface{}
	for _, p := range ports {
		ui = append(ui, p)
	}
	_, err = collection.InsertMany(context.Background(), ui)
	if err != nil {
		return err
	}
	return nil
}

func Get(start, limit int32) ([]*domain.Port, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("port").Collection("ports")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	ports := make([]*domain.Port, 0)
	for cur.Next(ctx) {
		var result *domain.Port
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		ports = append(ports, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return ports, nil
}
