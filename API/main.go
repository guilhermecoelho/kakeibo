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

	getRouter := routes.Methods(http.MethodGet).PathPrefix("/api").Subrouter()
	putRouter := routes.Methods(http.MethodPut).PathPrefix("/api").Subrouter()
	postRouter := routes.Methods(http.MethodPost).PathPrefix("/api").Subrouter()
	deleteRouter := routes.Methods(http.MethodDelete).PathPrefix("/api").Subrouter()

	postRouter.HandleFunc("/login/", handlers.Login)

	getRouter.HandleFunc("/group/", handlers.GetGroups)
	//getRouter.HandleFunc("/group/", configurations.IsAuthorized(handlers.GetGroups))

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

	getRouter.HandleFunc("/user/", handlers.GetUsers)
	getRouter.HandleFunc("/user/{id}", handlers.GetUserById)
	putRouter.HandleFunc("/user/", handlers.PutUser)
	postRouter.HandleFunc("/user/", handlers.PostUser)
	deleteRouter.HandleFunc("/user/", handlers.DeleteUser)

	log.Fatal(http.ListenAndServe(":8080", routes))
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}
