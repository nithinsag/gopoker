package helpers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/diadara/gopoker/user"
)

type key int

const tokenKey key = 0

// AuthenticationMiddleware that takes the authorization token and validates
// the token and puts it in context
// Rejects the request if invalid
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		fmt.Println("autj", auth)

		if len(auth) != 1 {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		tokenString := auth[0]
		isValid, token := user.Validate(string(tokenString))
		if !isValid {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), tokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
