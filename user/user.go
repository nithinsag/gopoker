package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/diadara/gopoker/db"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequestHandler returns a token if credentials are valid
func LoginRequestHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var credential Credential
	_ = dec.Decode(&credential)
	fmt.Println(credential)
	user, token := GetUserFromCredential(credential)
	//token := models.Token{Token: "sdafasdfasdf"}
	fmt.Println(user)
	json.NewEncoder(w).Encode(token)
}

// RequestHandler handles /me route
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You are logged in")
}

// User type
type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Password  string
}

// Credential type
type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token type
type Token struct {
	Token string `json:"token,omitempty"`
}

//MigrateUsers populates first users
func MigrateUsers() {
	fmt.Println(db.Db)
	db.Db.AutoMigrate(User{})
	//TODO: implement password hashing
	dbc := db.Db.Create(&User{Email: "nithin@fabelio.com", FirstName: "Maveli", Password: hashAndSalt([]byte("secretpassword"))})
	fmt.Println("created user")
	fmt.Println(dbc.Error)
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

func generateToken(user User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Email,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	fmt.Println("token", token)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tokenstring", tokenString)
	return tokenString
}

// Validate the tokenstring and return token if valid
func Validate(tokenString string) (bool, *jwt.Token) {
	// sample token string taken from the New example

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString("secret")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return true, token
	} else {
		fmt.Println(err)
	}
	return false, nil
}

// GetUserFromCredential gets the user from the credential
func GetUserFromCredential(credential Credential) (User, string) {
	var user User
	var token string

	fmt.Println(credential)
	db.Db.First(&user, "email = ?", credential.Email)
	fmt.Println("looked up user", user)
	fmt.Println("got credential", credential)
	authenticated := comparePasswords(user.Password, []byte(credential.Password))
	//TODO: implement password hashing
	fmt.Println(authenticated)
	if authenticated {
		token = generateToken(user)
	}
	fmt.Println("generated token", token)
	return user, token
}
