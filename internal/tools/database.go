package tools

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type Coinbalance struct {
	Username string
	Balance  int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *Coinbalance
	SetupDatabase() error
}

func NewDatabase(client *mongo.Client) (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{
		client: client,
	}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &database, nil

}
