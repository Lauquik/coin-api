package tools

import (
	"context"

	"github.com/lavquik/gorestapi/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mockDB struct {
	client *mongo.Client
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	var result db.UserData
	coll := d.client.Database("test").Collection("data")
	err := coll.FindOne(context.TODO(), bson.D{{"name", username}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return &LoginDetails{
		Username:  result.Creds.Username,
		AuthToken: result.Creds.AuthToken,
	}
}

func (d *mockDB) GetUserCoins(user string) *Coinbalance {
	var result db.UserData
	coll := d.client.Database("test").Collection("data")
	err := coll.FindOne(context.TODO(), bson.D{{"name", user}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return &Coinbalance{
		Username: result.Creds.Username,
		Balance:  int64(result.Balance.Balance),
	}
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
