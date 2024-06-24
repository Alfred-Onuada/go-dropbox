package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Alfred-Onuada/go-dropbox/internals/auth"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
      token := strings.Split(tokenString," ")

	  if len(token) != 2 {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
	  }
	  tokenString = token[1]
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to request context
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}