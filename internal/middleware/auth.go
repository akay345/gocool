
package middleware

import (
    "net/http"
    "strings"
    "github.com/dgrijalva/jwt-go"
)

// AuthMiddleware checks if the request is authenticated using JWT tokens.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenHeader := r.Header.Get("Authorization")
        if tokenHeader == "" {
            http.Error(w, "Missing auth token", http.StatusUnauthorized)
            return
        }

        // Typically the Authorization header will be in the format `Bearer {token}`
        // We check if the received header meets this requirement
        splitToken := strings.Split(tokenHeader, "Bearer ")
        if len(splitToken) != 2 {
            http.Error(w, "Invalid/Malformed auth token", http.StatusUnauthorized)
            return
        }

        jwtToken := splitToken[1]

        // Validate token
        token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
            // Make sure that the token method conform to "SigningMethodHMAC"
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.Error(w, "Unexpected signing method in auth token", http.StatusUnauthorized)
            }
            return []byte("YourSecretKey"), nil // Replace 'YourSecretKey' with your actual secret key
        })

        if err != nil {
            http.Error(w, "Invalid auth token", http.StatusUnauthorized)
            return
        }

        if !token.Valid {
            http.Error(w, "Invalid auth token", http.StatusUnauthorized)
            return
        }

        // Proceed to the next middleware or handler if authenticated
        next.ServeHTTP(w, r)
    })
}
