package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/diadara/gopoker/helpers/config"
	"github.com/diadara/gopoker/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type handler func(w http.ResponseWriter, r *http.Request)

func LoginRequestHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var credential models.Credential
	_ = dec.Decode(&credential)
	fmt.Println(credential)
	user, token := models.GetUserFromCredential(credential)
	//token := models.Token{Token: "sdafasdfasdf"}
	fmt.Println(user)
	json.NewEncoder(w).Encode(token)
}

func UserRequestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You are logged in")
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		fmt.Println("autj", auth)

		if len(auth) != 1 {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		tokenString := auth[0]
		isValid, token := models.Validate(string(tokenString))
		if !isValid {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// our main function
func main() {
	config.InitViper()

	router := mux.NewRouter()
	userRouter := router.PathPrefix("/me").Subrouter()
	userRouter.Use(authenticationMiddleware)
	models.InitDB()
	defer models.CloseDB()
	router.HandleFunc("/login", LoginRequestHandler).Methods("POST")
	userRouter.HandleFunc("", UserRequestHandler).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
