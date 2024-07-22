package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zinitdev/eSale/pkg/routes"
	"github.com/zinitdev/eSale/pkg/configs"
)

func main() {
	configs.LoadEnv()

	r := mux.NewRouter()
    routes.DefaultRoutes(r)
	http.Handle("/", r)

	port := configs.GetEnv("PORT")

	if port == "" {
        port = "8080"
    }

	fmt.Printf("Server is starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

