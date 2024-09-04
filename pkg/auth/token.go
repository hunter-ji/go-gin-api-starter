// @Title token.go
// @Description
// @Author Hunter 2024/9/4 16:48

package auth

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"go-gin-api-starter/config"
)

var (
	accessTokenSecret  []byte
	refreshTokenSecret []byte
	refreshMutexes     sync.Map
	tokenCache         sync.Map
)

type cachedToken struct {
	accessToken  string
	refreshToken string
	expiry       time.Time
}

func init() {
	accessTokenSecret = []byte(config.TokenConfig.AccessTokenSecret)
	refreshTokenSecret = []byte(config.TokenConfig.RefreshTokenSecret)
}

type Claims struct {
	UserID uint64 `json:"userID"`
	jwt.RegisteredClaims
}

// GenerateAccessToken
// @Description: Generate new access token
// @param userID uint64
// @return string access token
// @return error
func GenerateAccessToken(userID uint64) (string, error) {
	expirationTime := time.Now().Add(10 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecret)
}

func generateRefreshToken(userID uint64) (string, error) {
	expirationTime := time.Now().Add(15 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenSecret)
}

func ValidateAccessToken(tokenString string) (*Claims, bool, error) {
	return validateToken(tokenString, accessTokenSecret)
}

func validateRefreshToken(tokenString string) (*Claims, bool, error) {
	return validateToken(tokenString, refreshTokenSecret)
}

// validateToken
// @Description: validate token
// @param tokenString token string
// @param secret secret key
// @return *Claims
// @return bool expired
// @return error
func validateToken(tokenString string, secret []byte) (*Claims, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, true, err
		}
		return nil, false, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, false, nil
	}

	return nil, false, jwt.ErrSignatureInvalid
}

func storeRefreshToken(userID uint64, refreshToken string, redisClient *redis.Client) error {
	key := fmt.Sprintf("%s-refresh-token:%d", config.CommonSplicePrefix, userID)
	return redisClient.Set(context.Background(), key, refreshToken, 15*24*time.Hour).Err()
}

func validateStoredRefreshToken(userID uint64, refreshToken string, redisClient *redis.Client) bool {
	key := fmt.Sprintf("%s-refresh-token:%d", config.CommonSplicePrefix, userID)
	storedToken, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return storedToken == refreshToken
}

func deleteRefreshToken(userID uint64, redisClient *redis.Client) error {
	key := fmt.Sprintf("%s-refresh-token:%d", config.CommonSplicePrefix, userID)
	return redisClient.Del(context.Background(), key).Err()
}

// GenerateAccessTokenAndRefreshToken
// @Description: Generate new access token and refresh token, then store the refresh token
// @param userID uint64
// @param redisClient *redis.Client
// @return string access token
// @return string refresh token
// @return error
func GenerateAccessTokenAndRefreshToken(userID uint64, redisClient *redis.Client) (string, string, error) {
	// Generate new access token
	newAccessToken, err := GenerateAccessToken(userID)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	// Generate new refresh token
	newRefreshToken, err := generateRefreshToken(userID)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate new refresh token: %w", err)
	}

	// Store new refresh token
	if err := storeRefreshToken(userID, newRefreshToken, redisClient); err != nil {
		return "", "", fmt.Errorf("failed to store new refresh token: %w", err)
	}

	return newAccessToken, newRefreshToken, nil
}

// ValidateAccessTokenAndRefresh
// @Description: validate access token and refresh it if expired
// @param accessToken
// @param redisClient
// @return *Claims
// @return string new access token
// @return bool expired
// @return error
func ValidateAccessTokenAndRefresh(accessToken string, redisClient *redis.Client) (*Claims, string, bool, error) {
	// Validate access token
	claims, expired, err := ValidateAccessToken(accessToken)
	if err != nil {
		return nil, "", false, fmt.Errorf("invalid access token: %w", err)
	}

	// Return empty string if access token is not expired
	if !expired {
		return claims, "", false, nil
	}

	userID := claims.UserID

	// Check cache first
	if cachedToken, ok := getTokenFromCache(userID); ok {
		return claims, cachedToken.accessToken, false, nil
	}

	// Get or create a mutex for this userID
	mutex, _ := refreshMutexes.LoadOrStore(userID, &sync.Mutex{})
	mtx := mutex.(*sync.Mutex)
	mtx.Lock()
	defer mtx.Unlock()

	// Check cache again after acquiring the lock
	if cachedToken, ok := getTokenFromCache(userID); ok {
		return claims, cachedToken.accessToken, false, nil
	}

	// Get refresh token from redis
	key := fmt.Sprintf("%s-refresh-token:%d", config.CommonSplicePrefix, userID)
	refreshToken, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, "", true, fmt.Errorf("failed to get refresh token: %w", err)
	}

	// Validate refresh token
	if _, _, err := validateRefreshToken(refreshToken); err != nil {
		return nil, "", true, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Generate new tokens
	newAccessToken, _, err := GenerateAccessTokenAndRefreshToken(userID, redisClient)
	if err != nil {
		return nil, "", false, err
	}

	return claims, newAccessToken, false, nil
}

// RefreshToken
// @Description: Refresh access token and refresh token
// @param refreshToken string
// @param redisClient *redis.Client
// @return string access token
// @return string refresh token
// @return error
func RefreshToken(refreshToken string, redisClient *redis.Client) (string, string, error) {
	// Validate refresh token
	claims, expired, err := validateRefreshToken(refreshToken)
	if err != nil {
		if expired {
			return "", "", errors.New("refresh token expired")
		}
		return "", "", fmt.Errorf("invalid refresh token: %w", err)
	}

	userID := claims.UserID

	// Check cache first
	if cachedToken, ok := getTokenFromCache(userID); ok {
		return cachedToken.accessToken, cachedToken.refreshToken, nil
	}

	// Get or create a mutex for this userID
	mutex, _ := refreshMutexes.LoadOrStore(userID, &sync.Mutex{})
	mtx := mutex.(*sync.Mutex)
	mtx.Lock()
	defer mtx.Unlock()

	// Check cache again after acquiring the lock
	if cachedToken, ok := getTokenFromCache(userID); ok {
		return cachedToken.accessToken, cachedToken.refreshToken, nil
	}

	// Validate stored refresh token
	if !validateStoredRefreshToken(userID, refreshToken, redisClient) {
		return "", "", errors.New("invalid refresh token")
	}

	// Generate new tokens
	newAccessToken, newRefreshToken, err := GenerateAccessTokenAndRefreshToken(userID, redisClient)
	if err != nil {
		return "", "", err
	}

	// Cache the new tokens
	cacheToken(userID, newAccessToken, newRefreshToken)

	return newAccessToken, newRefreshToken, nil
}

// DeleteToken
// @Description: Delete token
// @param userID
// @param redisClient
// @return error
func DeleteToken(userID uint64, redisClient *redis.Client) error {
	if err := deleteRefreshToken(userID, redisClient); err != nil {
		return err
	}
	return nil
}

func getTokenFromCache(userID uint64) (cachedToken, bool) {
	if tokenInterface, ok := tokenCache.Load(userID); ok {
		token := tokenInterface.(cachedToken)
		if time.Now().Before(token.expiry) {
			return token, true
		}
		// Expired cache, remove it
		tokenCache.Delete(userID)
	}
	return cachedToken{}, false
}

func cacheToken(userID uint64, accessToken, refreshToken string) {
	tokenCache.Store(userID, cachedToken{
		accessToken:  accessToken,
		refreshToken: refreshToken,
		expiry:       time.Now().Add(5 * time.Minute), // cache for 5 minutes
	})
}
