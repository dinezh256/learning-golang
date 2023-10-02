package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	muxRouter := mux.NewRouter()
	routes.RegisterBookStoreRoutes(muxRouter)
	http.Handle("/", muxRouter)
	log.Fatal(http.ListenAndServe("localhost:8080", muxRouter))
}
