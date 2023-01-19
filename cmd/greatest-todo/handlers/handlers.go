package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/darwin808/greatest-todo/cmd/greatest-todo/models"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Accounts

	if result := h.DB.Find(&accounts); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var account models.Accounts
	json.Unmarshal(body, &account)

	// Append to the Books table
	if result := h.DB.Create(&account); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
