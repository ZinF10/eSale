package routes

import (
	"github.com/gorilla/mux"
	"github.com/zinitdev/eSale/pkg/controllers"
)

var DefaultRoutes = func(router *mux.Router) {
	router.HandleFunc("/categories/", controllers.GetCategories).Methods("GET")
}