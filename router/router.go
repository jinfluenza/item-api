package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinfluenza/item-api/handlers"
	model "github.com/jinfluenza/item-api/models"
	log "github.com/sirupsen/logrus"
)

// Get all router
func GetItemsRouter(w http.ResponseWriter, r *http.Request) {
	log.Info("Attempting to get the items")

	items := handlers.GetItems()

	w = addSuccessHeaders(w)

	json.NewEncoder(w).Encode(items)
}

func GetItemByTitleRouter(w http.ResponseWriter, r *http.Request) {
	log.Info("Attempting to get specific item by its title")

	title := r.URL.Query().Get("title")

	if title == "" {
		log.Errorf("Data was null, and failed due to this reason: title param query was not found")
	}

	it, err := handlers.GetItemByTitle(title)

	if err != nil {
		log.Errorf("Failed reason: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w = addSuccessHeaders(w)
		json.NewEncoder(w).Encode(it)
	}
}

func CreateItemRouter(w http.ResponseWriter, r *http.Request) {
	log.Info("Creating item now")

	var item model.Item

	rb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Errorf("Error due to following reason: %s", err)
	}

	json.Unmarshal(rb, &item)

	finalItem, err := handlers.CreateItems(item)

	if err != nil {
		log.Errorf("Error while processing the data: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w = addSuccessHeaders(w)
		json.NewEncoder(w).Encode(finalItem)
	}
}

func UpdateItemRouter(w http.ResponseWriter, r *http.Request) {
	log.Info("Updating the item!")

	var item model.Item

	rb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Errorf("Error due to following reason: %s", err)
	}

	json.Unmarshal(rb, &item)

	finalItem, err := handlers.UpdateItem(item.Title, item)

	if err != nil {
		log.Errorf("Error while processing the data: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w = addSuccessHeaders(w)
		json.NewEncoder(w).Encode(finalItem)
	}
}

func DeleteItemRouter(w http.ResponseWriter, r *http.Request) {
	log.Info("Deleting the item!")

	var item model.Item

	rb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Errorf("Error due to following reason: %s", err)
	}

	json.Unmarshal(rb, &item)

	_, err = handlers.DeleteItem(item)

	if err != nil {
		log.Errorf("Error while processing the data: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
	} else {
		w = addSuccessHeaders(w)
		json.NewEncoder(w).Encode("Item was deleted")
	}
}

func addSuccessHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return w
}
