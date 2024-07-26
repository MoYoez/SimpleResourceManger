package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
	utils "github.com/moyoez/starry-backend/utils"
)

/*

type UserInfoDataStruct struct {
	UserID        string `db:"user_id" json:"user_id"` // UNIQUE
	UserName      string `db:"user_name" json:"user_name"`
	UserPassword  string `db:"user_password" json:"user_password"` // Password was store in hasher.
	UserEmail     string `db:"user_email" json:"user_email"`
	UserAuthority int64  `db:"user_authority" json:"user_authority"` // 0 => root | 1 => admin | 2 => normal User.
	UserSex       int64  `db:"user_sex" json:"user_sex"`             // 0 ==> male | 1 ==> female
}

*/

func Register(c *gin.Context) {
	var RegisterBody sql.UserInfoDataStruct
	if c.BindJSON(&RegisterBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "BAD REQUEST",
		})
		return
	}
	getRole, _ := c.Get("user")
	getRoleNum := sql.QueryDataBaseByID(getRole.(string)).UserAuthority
	getRoleExisted := sql.QueryDataBaseByID(getRole.(string)).UserName
	if getRoleExisted == "" {
		getRoleNum = 3
	}
	if RegisterBody.UserAuthority < 2 && getRoleNum >= 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Only ROOT ADMIN can modify admin permission.",
		})
		return
	}
	if sql.QueryDataBaseByID(RegisterBody.UserID).UserID != "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Account has been registered.",
		})
		return
	}
	// Write into Database.
	getHasherPassword := utils.Hasher(RegisterBody.UserPassword)
	sql.ModifyReferData(sql.UserInfoDataStruct{
		UserID:        RegisterBody.UserID,
		UserName:      RegisterBody.UserName,
		UserPassword:  getHasherPassword,
		UserEmail:     RegisterBody.UserEmail,
		UserAuthority: RegisterBody.UserAuthority,
		UserSex:       RegisterBody.UserSex,
	})
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Register Success!",
	})
}
