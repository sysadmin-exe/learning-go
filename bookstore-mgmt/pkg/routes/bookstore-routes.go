package routes

import (
	"github.com/gorilla/mux"
	"github.com/sysadmin-exe/learning-go/bookstore-mgmt/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
