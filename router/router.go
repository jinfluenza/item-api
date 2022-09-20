package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinfluenza/item-api/handlers"
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

	var title string
	
	rb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("Data was null, and failed due to this reason: %s", err)
	}

	json.Unmarshal(rb, &title)

	it, err := handlers.GetItemByTitle(title)

	if err != nil {
		log.Errorf("Failed reason: %s", err)
		w = addFailedHeaders(w)
		json.NewEncoder(w).Encode(it)
	} else {
		w = addSuccessHeaders(w)
		json.NewEncoder(w).Encode(it)
	}

	
	

}

func addFailedHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.WriteHeader(http.StatusNotFound)
	return w
}

func addSuccessHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return w
}
