package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diadara/gopoker/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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

// our main function
func main() {
	router := mux.NewRouter()
	models.InitDB()
	defer models.CloseDB()

	router.HandleFunc("/login", LoginRequestHandler).Methods("POST")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
