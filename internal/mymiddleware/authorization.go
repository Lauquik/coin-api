package mymiddleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/lavquik/gorestapi/cmd/api"
	"github.com/lavquik/gorestapi/internal/tools"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrFooUnAuthorized = errors.New("invalid username or token")

type Authorzie struct {
	Client *mongo.Client
}

func NewAuth(client *mongo.Client) *Authorzie {
	return &Authorzie{
		Client: client,
	}
}

func (a *Authorzie) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		if username == "" || token == "" {
			api.RequestError(w, ErrFooUnAuthorized)
			return
		}
		var database *tools.DatabaseInterface
		database, err := tools.NewDatabase(a.Client)
		if err != nil {
			api.InternalErrorHander(w)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)
		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Fatal(ErrFooUnAuthorized)
			api.RequestError(w, ErrFooUnAuthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
