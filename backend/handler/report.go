package handler

import (
	"github.com/moyoez/starry-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/sql"
)

/*
	ReportID              string `db:"report_id" json:"report_id"`
	ReportContent         string `db:"report_content" json:"report_content"`
	ReportTime            int64  `db:"report_time" json:"report_time"`
	ReportIsEmergency     bool   `db:"report_emergency" json:"report_emergency"`
	ReportProcessStatus   int64  `db:"report_process" json:"report_process"` // 0 => no start 1 => started 2 => done.
	ReportFeedbackTime    int64  `db:"report_feedback_time" json:"report_feedback_time"`
	ReportFeedbackRate    int64  `db:"report_feedback_rate" json:"report_feedback_rate"` // feedback Rate 0-5
	ReportFeedbackContent string `db:"report_feedback_content" json:"report_feedback_content"`
	ReportFeedbackLogs    string `db:"report_feedback_logs" json:"report_feedback_logs"`
	ReportFeedbackUser    string `db:"report_feedback_user" json:"report_feedback_user"`

*/

func QueryFullReportList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success! ", "list": sql.QueryAllReportInfoData()})
}

func AddReport(c *gin.Context) {
	var resBind sql.ReportBaseStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	for _, data := range sql.QueryAllReportInfoData() {
		if data.ReportID == resBind.ReportID {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false, "message": "ID 不可重复! ",
			})
			return
		}
	}
	middleware.RequireAuth(c)
	getRole, _ := c.Get("user")
	getRoleNum := sql.QueryDataBaseByID(getRole.(string)).UserAuthority
	getRoleExisted := sql.QueryDataBaseByID(getRole.(string)).UserName
	if getRoleExisted == "" {
		getRoleNum = 3
	}

	if getRoleNum > 2 && resBind.ReportUser != getRoleExisted {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Normal User can only modify their own data.",
		})
		return
	}
	// query modifier data
	sql.ModifyReportReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Add Success!"})
}

func ModifyReport(c *gin.Context) {
	var resBind sql.ReportBaseStruct
	err := c.BindJSON(&resBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "Bad Request",
		})
		return
	}
	if sql.QueryReportDataBaseByID(resBind.ReportID).ReportID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "message": "未找到此资源",
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
	if getRoleNum > 2 && resBind.ReportUser != getRoleExisted {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Normal User can only modify their own data.",
		})
		return
	}
	sql.ModifyReportReferData(resBind)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Modify Success!"})
}

func DeleteReport(c *gin.Context) {
	var resBind sql.ReportBaseStruct
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
	if getRoleNum >= 2 && sql.QueryReportDataBaseByID(resBind.ReportID).ReportUser != getRoleExisted {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": false, "message": "Normal User can only modify their own data.",
		})
		return
	}
	sql.RemoveReferReportInfo(resBind.ReportID)
	c.JSON(http.StatusOK, gin.H{
		"Status": true, "message": "Delete Success!",
	})
}
