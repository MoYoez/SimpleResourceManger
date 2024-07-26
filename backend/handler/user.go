package handler

import (
	"github.com/moyoez/starry-backend/middleware"
	utils "github.com/moyoez/starry-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
)

func QueryUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Query list success.", "list": sql.QueryAllUserInfoDataWithNoPassword()})
}

func ModifyUserList(c *gin.Context) {
	// always can modify one.
	var resBind sql.UserInfoDataStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	getRole, _ := c.Get("user")
	getRoleNum := sql.QueryDataBaseByID(getRole.(string)).UserAuthority
	getRoleExisted := sql.QueryDataBaseByID(getRole.(string)).UserName
	if getRoleExisted == "" {
		getRoleNum = 3
	}
	getTargetRole := sql.QueryDataBaseByID(resBind.UserID).UserAuthority

	if getTargetRole < 2 && getRoleNum >= 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Only ROOT ADMIN can modify admin permission.",
		})
		return
	}

	sql.ModifyReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Modify Success!"})
}

func DelUserList(c *gin.Context) {
	// always can modify one.
	var resBind sql.UserInfoDataStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	middleware.RequireAuth(c)
	getRole, _ := c.Get("user")
	getRoleNum := sql.QueryDataBaseByID(getRole.(string)).UserAuthority
	getRoleExisted := sql.QueryDataBaseByID(getRole.(string)).UserName
	if getRoleExisted == "" {
		getRoleNum = 3
	}
	getTargetRole := sql.QueryDataBaseByID(resBind.UserID).UserAuthority

	if getTargetRole < 2 && getRoleNum >= 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Only ROOT ADMIN can modify admin permission.",
		})
		return
	}

	sql.RemoveReferUserInfo(resBind.UserID)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Delete Success!"})
}

func QuerySelfUserInfo(c *gin.Context) {
	getRole, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized User."})
		return
	}
	getRoleinfo := sql.QueryDataBaseByID(getRole.(string))
	c.JSON(http.StatusOK, getRoleinfo)
}

func ModifySelfUserInfo(c *gin.Context) {
	var resBind sql.UserInfoDataStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Bad Request"})
		return
	}
	middleware.RequireAuth(c)
	getRole, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized User."})
		return
	}
	if resBind.UserPassword == "" {
		sql.ModifyReferData(sql.UserInfoDataStruct{
			UserID:        getRole.(string),
			UserName:      resBind.UserName,
			UserPassword:  sql.QueryDataBaseByID(getRole.(string)).UserPassword,
			UserEmail:     resBind.UserEmail,
			UserAuthority: resBind.UserAuthority,
			UserSex:       resBind.UserSex,
		})
	} else {
		sql.ModifyReferData(sql.UserInfoDataStruct{
			UserID:        getRole.(string),
			UserName:      resBind.UserName,
			UserPassword:  utils.Hasher(resBind.UserPassword),
			UserEmail:     resBind.UserEmail,
			UserAuthority: resBind.UserAuthority,
			UserSex:       resBind.UserSex,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true, "message": "Update Success.",
	})

}
