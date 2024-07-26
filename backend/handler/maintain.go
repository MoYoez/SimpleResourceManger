package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
)

/*
type MaintainStruct struct {
	MainTainID      int64  `db:"maintain_id"`
	MainTainChecker bool   `db:"maintain_checker"` // true == checked.
	MaintainTime    int64  `db:"maintain_time"`    // use uinx 64
	MaintainContent string `db:"maintain_content"` // maintain content
}
*/

func QueryMaintainList(c *gin.Context) {
	getList := sql.QueryAllMaintainInfoData()
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "is ok! ", "list": getList})
}

func AddMaintainList(c *gin.Context) {
	// always can modify one.
	var resBind sql.MaintainTableStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	for _, data := range sql.QueryAllMaintainInfoData() {
		if data.MaintainID == resBind.MaintainID {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false, "message": "ID 不可重复! ",
			})
			return
		}
	}
	sql.ModifyMaintainReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Add Success!"})
}

func ModifyMaintainList(c *gin.Context) {
	// always can modify one.
	var resBind sql.MaintainTableStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	if sql.QueryMaintainDataBaseByID(resBind.MaintainID).MaintainID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "未找到此资源",
		})
		return
	}
	sql.ModifyMaintainReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Modify Success!"})
}

func DelMaintainList(c *gin.Context) {
	var resBind sql.MaintainTableStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	sql.RemoveReferMaintainInfo(resBind.MaintainID)
	c.JSON(http.StatusOK, gin.H{
		"Status":  true,
		"message": "Delete Success!",
	})
}
