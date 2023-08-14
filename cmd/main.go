package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lavquik/gorestapi/internal/db"
	"github.com/lavquik/gorestapi/internal/handler"
	"github.com/lavquik/gorestapi/internal/mymiddleware"
)

func main() {
	var r *chi.Mux = chi.NewRouter()
	client := db.Newmongoclient()
	auth := mymiddleware.NewAuth(client)
	h := handler.NewHandler(client)
	h.Handler(r, auth)
	fmt.Println("starting go api service.....")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
