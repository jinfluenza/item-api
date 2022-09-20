package handlers

import (
	"fmt"

	model "github.com/jinfluenza/item-api/models"
	log "github.com/sirupsen/logrus"
)

var fancy_pants_db []model.Item

// METHOD:GET
func GetItemByTitle(title string) (model.Item, error) {
	var item model.Item

	for _, v := range fancy_pants_db {
		if v.Title == title {
			item = v
		}
	}
	if item.Title == "" {
		log.Fatal("Item not found :(")
		return item, fmt.Errorf("Item not found :(")
	}
	log.Info("Title %s was found", title)
	return item, nil
}

// METHOD:GETALL
func GetItems() []model.Item {
	item := model.Item{Title: "hello", Body: "Bye"}
	fancy_pants_db = append(fancy_pants_db, item)
	return fancy_pants_db
}

// METHOD:CREATE
func CreateItems(i model.Item) (model.Item, error) {
	if i.Title == "" || i.Body == "" {
		return i, fmt.Errorf("Please send correct data!")
	}
	fancy_pants_db = append(fancy_pants_db, i)
	return i, nil
}

// METHOD:UPDATE
func UpdateItem(title string, ei model.Item) (model.Item, error) {
	var editedItem model.Item

	for i, v := range fancy_pants_db {
		if v.Title == title {
			fancy_pants_db[i] = ei
			editedItem = ei
		}
	}

	if editedItem.Title == "" {
		log.Fatal("Item not found")
		return editedItem, fmt.Errorf("Item was not found in the db!")
	} else {
		log.Info("%s was found in the db and updated!", title)
	}

	return editedItem, nil
}

// METHOD:DELTE
func DeleteItem(item model.Item) (model.Item, error) {
	var deletedItem model.Item

	for i, v := range fancy_pants_db {
		if v.Title == item.Title && v.Body == item.Body {
			fancy_pants_db = append(fancy_pants_db[:i], fancy_pants_db[i+1:]...)
			deletedItem = item
			break
		}
	}

	if deletedItem.Title == "" {
		return deletedItem, fmt.Errorf("No matching data to delete!!")
	}

	return deletedItem, nil
}
