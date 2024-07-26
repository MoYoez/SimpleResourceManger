package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
	"net/http"
)

/*

type ResouceMangerStruct struct {
	ResID    int64  `db:"res_id"`    // Res ID IS ONLY ONE.
	ResCount int64  `db:"res_count"` // Res Count is here.
	ResType  string `db:"res_type"`
	ResModel string `db:"res_model"`
}
*/

func ResMangerGet(c *gin.Context) {
	/*
		var resBind sql.ResouceMangerStruct
		err := c.BindJSON(&resBind)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Bad Request",
			})
			return
		}
	*/
	getResManagerList := sql.QueryAllResInfoData()
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Get list Success", "list": getResManagerList})
}

func ResManagerAdd(c *gin.Context) {
	// always can modify one.
	var resBind sql.ResBaseStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	for _, data := range sql.QueryAllResInfoData() {
		if data.ResID == resBind.ResID {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "ID 不可重复! ",
			})
			return
		}
	}
	sql.ModifyResReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Add Success!"})
}

// ResMangerModify ResManager Cannot modify ResID, as is intended to be the final.
func ResMangerModify(c *gin.Context) {
	// always can modify one.
	var resBind sql.ResBaseStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	if sql.QueryResDataBaseByID(resBind.ResID).ResID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "未找到此资源",
		})
		return
	}
	sql.ModifyResReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Modify Success!"})
}

func ResmanagerDelete(c *gin.Context) {
	var resBind sql.ResBaseStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	sql.RemoveReferResInfo(resBind.ResID)
	c.JSON(http.StatusOK, gin.H{
		"Status": true, "message": "Delete Success!",
	})
}
