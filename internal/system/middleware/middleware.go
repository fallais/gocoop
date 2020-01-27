package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"gocoop/internal/cache"
	"gocoop/internal/protocols"

	"github.com/alioygur/gores"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

//------------------------------------------------------------------------------
// Middleware
//------------------------------------------------------------------------------

// JWT is a middleware used to check JSON Web Token
func JWT(store *cache.RedisCache, privateKey string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Read the Authorization header
			header := r.Header.Get("Authorization")
			if header == "" {
				response := &protocols.APIControllerResponse{
					ErrorID:      "MIDDLEWARE/JWT/001",
					ErrorMessage: "Missing authorization header",
				}

				gores.JSON(w, http.StatusUnauthorized, response)
				return
			}

			// Split header into two parts
			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				response := &protocols.APIControllerResponse{
					ErrorID:      "MIDDLEWARE/JWT/002",
					ErrorMessage: "Invalid authorization header",
				}

				gores.JSON(w, http.StatusUnauthorized, response)
				return
			}

			// Get the token
			tokenRaw := headerParts[1]

			// Parse the token
			token, err := jwt.Parse(tokenRaw, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(privateKey), nil
			})

			if err != nil {
				response := &protocols.APIControllerResponse{
					ErrorID:      "MIDDLEWARE/JWT/003",
					ErrorMessage: fmt.Sprintf("Unable to parse the token : %s", err),
				}

				gores.JSON(w, http.StatusUnauthorized, response)
				return
			}

			// Get the token from the database
			_, err = store.Get(tokenRaw)
			if err != nil {
				response := &protocols.APIControllerResponse{
					ErrorID:      "MIDDLEWARE/JWT/004",
					ErrorMessage: "Token does not exist in cache",
				}

				gores.JSON(w, http.StatusUnauthorized, response)
				return
			}

			// Prepare the context
			tokenParsed := token.Claims.(jwt.MapClaims)
			ctx := context.WithValue(r.Context(), "token", tokenRaw)
			ctx = context.WithValue(ctx, "email", tokenParsed["email"])
			ctx = context.WithValue(ctx, "role", tokenParsed["role"])
			ctx = context.WithValue(ctx, "_id", tokenParsed["_id"])

			// Serve
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

// Cors is a Goji middleware to handle CORS (from the great Zenithar)
func Cors() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			// Stop here if its Preflighted OPTIONS request
			if r.Method == "OPTIONS" {
				return
			}

			// Serve
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
