package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses ditolak. Header Authorization tidak ditemukan."})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "fallback-secret-sementara"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode penandatanganan tidak valid")
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau telah kedaluwarsa"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			var userID interface{}
			if claims["sub"] != nil {
				userID = claims["sub"]
			} else if claims["user_id"] != nil {
				userID = claims["user_id"]
			}

			if userID == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Identitas pengguna tidak ditemukan dalam token"})
				c.Abort()
				return
			}

			c.Set("userID", userID)
			c.Set("userRole", claims["role"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim token tidak valid"})
			c.Abort()
		}
	}
}
