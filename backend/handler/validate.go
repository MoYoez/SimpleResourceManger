package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
	"net/http"
)

func Validate(c *gin.Context) {
	getUser, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "done!", "user": getUser, "auth": sql.QueryDataBaseByID(getUser.(string)).UserAuthority})
}
