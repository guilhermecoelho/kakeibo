package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/handlers"
)

func main() {

	go configurations.InitDatabaseGorm()

	routes := mux.NewRouter().StrictSlash(true)

	getRouter := routes.Methods(http.MethodGet).Subrouter()
	putRouter := routes.Methods(http.MethodPut).Subrouter()
	postRouter := routes.Methods(http.MethodPost).Subrouter()
	deleteRouter := routes.Methods(http.MethodDelete).Subrouter()

	getRouter.HandleFunc("/group/", handlers.GetGroups)
	getRouter.HandleFunc("/group/{id}", handlers.GetGroupById)
	putRouter.HandleFunc("/group/", handlers.PutGroup)
	postRouter.HandleFunc("/group/", handlers.PostGroup)
	deleteRouter.HandleFunc("/group/", handlers.DeleteGroup)

	log.Fatal(http.ListenAndServe(":8080", routes))
}
