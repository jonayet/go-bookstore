package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jonayet/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Listening to 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
