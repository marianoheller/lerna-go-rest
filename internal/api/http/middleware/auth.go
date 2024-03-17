package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
)

func protectedRoute() func(http.ResponseWriter, *http.Request) {
	clerk.SetKey("​sk_test_••••••••••••••••••••••••••••••••••​")

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session JWT from the Authorization header
		sessionToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		// Verify the session
		claims, err := jwt.Verify(r.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {
			// handle the error
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"access": "unauthorized"}`))
			return
		}
		fmt.Fprintf(w, `{"user_id": "%s"}`, claims.Subject)
	}
}
