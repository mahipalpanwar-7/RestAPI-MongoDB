package router

import (
	"github.com/gorilla/mux"
	"github.com/mahipalpanwar-7/RestAPI-MongoDB/controller"
)

func Router() *mux.Router{
	router := mux.NewRouter()
    router.HandleFunc("/api/movies",controller.GetMYAllMovies).Methods("GET")
    router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
    router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
    router.HandleFunc("/api/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
    router.HandleFunc("/api/deleteallmovie",controller.DeleteAllMovie).Methods("DELETE")


	return router 
}