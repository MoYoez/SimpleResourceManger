package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
	"net/http"
)

func GetUserRole(c *gin.Context) int64 {
	getRole, ok := c.Get("user")
	if !ok {
		return 3 // GUEST
	}
	getRoleNum := sql.QueryDataBaseByID(getRole.(string)).UserAuthority
	getRoleExisted := sql.QueryDataBaseByID(getRole.(string)).UserID
	if getRoleExisted == "" {
		return 3
	}
	return getRoleNum
}

func IsCanModifiyDataUser(c *gin.Context) {
	if GetUserRole(c) < 2 {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "No authority to modify data."})
	}
}
func IsCanSetAdminUser(c *gin.Context) {
	if GetUserRole(c) == 0 {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "No authority to set admin."})
	}
}

func IsCanModifyMaintainSure(c *gin.Context) {
	if GetUserRole(c) == 0 {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "No authority to modify Verify."})
	}
}
