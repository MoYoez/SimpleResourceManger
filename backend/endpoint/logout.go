package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Logout Success!"})
}
