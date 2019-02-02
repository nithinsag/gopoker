package main

import (
	"log"
	"net/http"

	"github.com/diadara/gopoker/db"
	"github.com/diadara/gopoker/helpers"
	"github.com/diadara/gopoker/helpers/config"
	"github.com/diadara/gopoker/user"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type handler func(w http.ResponseWriter, r *http.Request)

// our main function
func main() {
	config.InitViper()

	router := mux.NewRouter()
	userRouter := router.PathPrefix("/me").Subrouter()
	userRouter.Use(helpers.AuthenticationMiddleware)
	db.InitDB()
	defer db.CloseDB()
	router.HandleFunc("/login", user.LoginRequestHandler).Methods("POST")
	userRouter.HandleFunc("", user.RequestHandler).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
