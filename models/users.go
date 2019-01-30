package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Password  string
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token,omitempty"`
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	fmt.Println(string(hash))
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	fmt.Println(hashedPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func MigrateUsers() {
	fmt.Println(db)
	db.AutoMigrate(User{})
	//TODO: implement password hashing
	dbc := db.Create(&User{Email: "nithin@fabelio.com", FirstName: "Maveli", Password: hashAndSalt([]byte("secretpassword"))})
	fmt.Println("created user")
	fmt.Println(dbc.Error)
}

func GetUserFromCredential(credential Credential) User {
	var user User
	fmt.Println(credential)
	db.First(&user, "email = ?", credential.Email)
	fmt.Println("looked up user", user)
	fmt.Println("got credential", credential)
	authenticated := comparePasswords(user.Password, []byte(credential.Password))
	//TODO: implement password hashing
	fmt.Println(authenticated)
	return user
}
