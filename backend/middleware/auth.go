package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/moyoez/starry-backend/sql"
	"net/http"
	"time"
)

// RequireAuth if no login, return to login page.
func RequireAuth(c *gin.Context) {
	getCookie := c.GetHeader("Authorization")
	if getCookie == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "No login."})
		return
	}
	token, _ := jwt.Parse(getCookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("sign Method error")
		}
		return []byte("SECRET"), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Token expired."})
			return
		}
		getUserInterface := sql.QueryDataBaseByID(claims["sub"].(string))
		if getUserInterface.UserID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "No user."})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Token error."})
		return
	}
	c.Set("user", claims["sub"])
	c.Next()
}
