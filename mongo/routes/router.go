package routes

import (
	"github.com/gorilla/mux"
	"github.com/hydraveer/workingwithmongo/controllers"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/movies",controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movie",controllers.CreateMoview).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}",controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}",controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/v1/deleteallmovie",controllers.DeleteAllVideo).Methods("DELETE")
	return router
}