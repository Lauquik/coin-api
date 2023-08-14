package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type insert struct {
	Client *mongo.Client
}

func NewInsert(client *mongo.Client) *insert {
	return &insert{
		Client: client,
	}
}

func (i *insert) InserInDb(userdata *UserData) *error {
	coll := i.Client.Database("test").Collection("data")
	result, err := coll.InsertOne(context.TODO(), &userdata)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return &err
}
