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

	getRouter.HandleFunc("/group/", configurations.IsAuthorized(handlers.GetGroups))

	getRouter.HandleFunc("/group/{id}", configurations.IsAuthorized(handlers.GetGroupById))
	putRouter.HandleFunc("/group/", configurations.IsAuthorized(handlers.PutGroup))
	postRouter.HandleFunc("/group/", configurations.IsAuthorized(handlers.PostGroup))
	deleteRouter.HandleFunc("/group/{id}", configurations.IsAuthorized(handlers.DeleteGroup))

	getRouter.HandleFunc("/category/", configurations.IsAuthorized(handlers.GetCategories))
	getRouter.HandleFunc("/category/{id}", configurations.IsAuthorized(handlers.GetCategoryById))
	putRouter.HandleFunc("/category/", configurations.IsAuthorized(handlers.PutCategory))
	postRouter.HandleFunc("/category/", configurations.IsAuthorized(handlers.PostCategory))
	deleteRouter.HandleFunc("/category/{id}", configurations.IsAuthorized(handlers.DeleteCategory))

	getRouter.HandleFunc("/expense/", configurations.IsAuthorized(handlers.GetExpenses))
	getRouter.HandleFunc("/expense/{id}", configurations.IsAuthorized(handlers.GetExpenseById))
	putRouter.HandleFunc("/expense/", configurations.IsAuthorized(handlers.PutExpense))
	postRouter.HandleFunc("/expense/", configurations.IsAuthorized(handlers.PostExpense))
	deleteRouter.HandleFunc("/expense/{id}", configurations.IsAuthorized(handlers.DeleteExpense))

	getRouter.HandleFunc("/income/", configurations.IsAuthorized(handlers.GetIncomes))
	getRouter.HandleFunc("/income/{id}", configurations.IsAuthorized(handlers.GetIncomeById))
	putRouter.HandleFunc("/income/", configurations.IsAuthorized(handlers.PutIncome))
	postRouter.HandleFunc("/income/", configurations.IsAuthorized(handlers.PostIncome))
	deleteRouter.HandleFunc("/income/{id}", configurations.IsAuthorized(handlers.DeleteIncome))

	getRouter.HandleFunc("/user/", configurations.IsAuthorized(handlers.GetUsers))
	getRouter.HandleFunc("/user/{id}", configurations.IsAuthorized(handlers.GetUserById))
	getRouter.HandleFunc("/user/name/{name}", configurations.IsAuthorized(handlers.GetUserByName))
	putRouter.HandleFunc("/user/", configurations.IsAuthorized(handlers.PutUser))
	postRouter.HandleFunc("/user/", configurations.IsAuthorized(handlers.PostUser))
	deleteRouter.HandleFunc("/user/{id}", configurations.IsAuthorized(handlers.DeleteUser))

	postRouter.HandleFunc("/report/", configurations.IsAuthorized(handlers.GetReport))

	log.Fatal(http.ListenAndServe(":8080", routes))
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}
