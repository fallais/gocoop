package services

import (
	"time"

	"github.com/fallais/gocoop/internal/cache"

	"github.com/dgrijalva/jwt-go"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type jwtService struct {
	store                *cache.RedisCache
	privateKey           string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// Claims is the struct.
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewJwtService returns a new JWT service
func NewJwtService(store *cache.RedisCache, privateKey string) JwtService {
	return &jwtService{
		store:                store,
		privateKey:           privateKey,
		accessTokenDuration:  12 * time.Hour,
		refreshTokenDuration: 24 * time.Hour,
	}
}

//------------------------------------------------------------------------------
// Services
//------------------------------------------------------------------------------

// Create a token
func (s *jwtService) Create(username string) (map[string]string, error) {
	// Create the claims
	accessTokenClaims := Claims{
		Username: username,
		Role:     "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.accessTokenDuration).Unix(),
			Issuer:    "watcher",
		},
	}

	// Create the token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	// Sign and get the complete encoded token as a string
	accessTokenString, err := accessToken.SignedString([]byte(s.privateKey))
	if err != nil {
		return nil, err
	}

	// Add token to the cache
	err = s.store.Set(accessTokenString, username, s.accessTokenDuration)
	if err != nil {
		return nil, err
	}

	// Create the claims
	refreshTokenClaims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.refreshTokenDuration).Unix(),
			Issuer:    "watcher",
		},
	}

	// Create the token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	// Sign and get the complete encoded token as a string
	refreshTokenString, err := refreshToken.SignedString([]byte(s.privateKey))
	if err != nil {
		return nil, err
	}

	// Add token to the cache
	err = s.store.Set(refreshTokenString, username, s.refreshTokenDuration)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

// Get a token
func (s *jwtService) Get(email string) (string, error) {
	return s.store.Get(email)
}
