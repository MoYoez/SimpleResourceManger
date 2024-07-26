package handler

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/moyoez/starry-backend/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LogininStruct struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

var SECRET = "114514"

func Login(c *gin.Context) {
	var mainBody LogininStruct
	err := c.BindJSON(&mainBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "bad request status",
		})
		return
	}
	// getLogin LogininStruct
	getHasher := sql.QueryDataBaseByID(mainBody.UserName).UserPassword

	if getHasher == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Invalid userName or UserPassword",
		})
		return
	}
	getBytes, err := hex.DecodeString(getHasher)
	if err != nil {
		panic(err)
	}
	if bcrypt.CompareHashAndPassword(getBytes, []byte(mainBody.UserPassword)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Invalid userName or UserPassword",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": mainBody.UserName,
		"exp": time.Now().Add(time.Hour * 3600).Unix(),
	})
	tokenstring, err := token.SignedString([]byte(SECRET))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Invalid to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Login Success!", "auth": tokenstring})
}
