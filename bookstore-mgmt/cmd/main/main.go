package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sysadmin-exe/learning-go/bookstore-mgmt/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	log.Println("Server started on: http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
