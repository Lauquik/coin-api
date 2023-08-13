package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/schema"
	"github.com/lavquik/gorestapi/cmd/api"
	"github.com/lavquik/gorestapi/internal/mymiddleware"
	"github.com/lavquik/gorestapi/internal/tools"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Client *mongo.Client
}

func NewHandler(client *mongo.Client) *Handler {
	return &Handler{
		Client: client,
	}
}

func (h *Handler) Handler(r *chi.Mux, a *mymiddleware.Authorzie) {

	r.Use(middleware.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(a.Authorization)
		router.Get("/coins", h.GetUserCoins)
	})
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hellow world")
	// })
}

func (h *Handler) GetUserCoins(w http.ResponseWriter, r *http.Request) {
	var params api.Checkbalacne
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Fatal(err)
		api.RequestError(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, errr := tools.NewDatabase(h.Client)
	if errr != nil {
		api.InternalErrorHander(w)
		return
	}

	balance := (*database).GetUserCoins(params.Username)
	if balance == nil {
		api.InternalErrorHander(w)
		return
	}

	var respose = api.CheckbalacneResp{
		Balance:    balance.Balance,
		StatusCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(respose)
	if err != nil {
		log.Fatal(err)
		api.InternalErrorHander(w)
		return
	}

}
