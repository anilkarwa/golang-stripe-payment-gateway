package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"nozlapp.com/modules/routes"
)

func main() {

	// setup router
	router := mux.NewRouter()

	// add all route files here
	routes.Router(router)

	// cors handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
	handler := cors.Default().Handler(router)

	// start http server
	fmt.Println("***** Server starting on port: 4001 ******")
	if err := http.ListenAndServe(":4001", handler); err != nil {
		log.Fatal(err)
	}
}
