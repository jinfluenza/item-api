package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	router "github.com/jinfluenza/item-api/router"
	log "github.com/sirupsen/logrus"
)

func requestHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/items", router.GetItemsRouter).Methods("GET")
	r.HandleFunc("/item", router.GetItemByTitleRouter).Methods("GET")
	r.HandleFunc("/item", router.CreateItemRouter).Methods("POST")
	r.HandleFunc("/item", router.UpdateItemRouter).Methods("PUT")
	r.HandleFunc("/item", router.DeleteItemRouter).Methods("DELETE")

	r.Use(cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
			AllowedHeaders: []string{"Accept", "content-type", "Content-Length"},
		}).Handler)

	log.Infoln("Serving the website on 4040")
	log.Errorln(http.ListenAndServe(":4040", handlers.CORS()(r)))
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	requestHandlers()
}
