package router

import "github.com/gorilla/mux"

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.Get)
}
