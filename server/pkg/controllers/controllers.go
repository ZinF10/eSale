package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/zinitdev/eSale/pkg/models"
)

var newCategories models.Category

func GetCategories(w http.ResponseWriter, r *http.Request) {
	newCategories := models.GetCategories()
	res, _ := json.Marshal(newCategories)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}