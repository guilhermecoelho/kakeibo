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

	getRouter.HandleFunc("/category/", handlers.GetCategories)
	getRouter.HandleFunc("/category/{id}", handlers.GetCategoryById)
	putRouter.HandleFunc("/category/", handlers.PutCategory)
	postRouter.HandleFunc("/category/", handlers.PostCategory)
	deleteRouter.HandleFunc("/category/", handlers.DeleteCategory)

	getRouter.HandleFunc("/expense/", handlers.GetExpenses)
	getRouter.HandleFunc("/expense/{id}", handlers.GetExpenseById)
	putRouter.HandleFunc("/expense/", handlers.PutExpense)
	postRouter.HandleFunc("/expense/", handlers.PostExpense)
	deleteRouter.HandleFunc("/expense/", handlers.DeleteExpense)

	getRouter.HandleFunc("/income/", handlers.GetIncomes)
	getRouter.HandleFunc("/income/{id}", handlers.GetIncomeById)
	putRouter.HandleFunc("/income/", handlers.PutIncome)
	postRouter.HandleFunc("/income/", handlers.PostIncome)
	deleteRouter.HandleFunc("/income/", handlers.DeleteIncome)

	log.Fatal(http.ListenAndServe(":8080", routes))
}
