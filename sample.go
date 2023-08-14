package main

import (
	"github.com/lavquik/gorestapi/internal/db"
)

var mockData []db.UserData = []db.UserData{
	{
		User: "alex",
		Creds: db.Credientials{
			Username:  "alex",
			AuthToken: "1234",
		},
		Balance: db.CoinBalance{
			Username: "alex",
			Balance:  10000,
		},
	},
	{
		User: "shubham",
		Creds: db.Credientials{
			Username:  "shubham",
			AuthToken: "111",
		},
		Balance: db.CoinBalance{
			Username: "shubham",
			Balance:  200000,
		},
	},
	{
		User: "laukik",
		Creds: db.Credientials{
			Username:  "laukik",
			AuthToken: "9900",
		},
		Balance: db.CoinBalance{
			Username: "laukik",
			Balance:  300000,
		},
	},
}

func main() {
	client := db.Newmongoclient()
	insert := db.NewInsert(client)

	for _, v := range mockData {
		insert.InserInDb(&v)
	}
	// var database *tools.DatabaseInterface
	// database, err := tools.NewDatabase(client)
	// if err != nil {
	// 	panic(err)
	// }
	// userbalance := (*database).GetUserCoins("laukik")
	// fmt.Println(userbalance)

}
