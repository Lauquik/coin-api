package db

type UserData struct {
	User    string       `bson:"name"`
	Creds   Credientials `bsons:"creds"`
	Balance CoinBalance  `bson:"balance"`
}

type Credientials struct {
	Username  string `bson:"username"`
	AuthToken string `bson:"authtoken"`
}

type CoinBalance struct {
	Username string `bson:"username"`
	Balance  int    `bson:"balance"`
}
