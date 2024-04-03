package routes

import (
	"log"
	"net/http"

	"github.com/danieldwipaska/muxserver/src/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomeHandler)
	router.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	router.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
