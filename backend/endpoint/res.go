package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/moyoez/starry-backend/handler"
	"github.com/moyoez/starry-backend/middleware"
)

func init() {
	SetRoot := R.Group("/backend", Cors())
	{
		// RES: ROLE USE MANAGER
		SetRoot.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "ALIVE")
		})
	}

	AnnoyGroup := SetRoot.Group("/annoy").Use(Cors())
	{
		/*
			AnnoyGroup.GET("/test", func(context *gin.Context) {
				context.String(200, "ALIVEss")
			})
		*/
		AnnoyGroup.POST("/login", handler.Login)
		AnnoyGroup.POST("/logout", Logout)
	}

	getNormalUserGroupManger := SetRoot.Group("/normal").Use(Cors(), middleware.RequireAuth)
	// Normal Manager means Group 0 | Group 1 | Group 2 ==> if people login their accounts , then he / she can check for a visit.
	{
		getNormalUserGroupManger.GET("/whoami", handler.Validate)
		getNormalUserGroupManger.GET("/res", handler.ResMangerGet)
		getNormalUserGroupManger.GET("/report", handler.QueryFullReportList)
		getNormalUserGroupManger.GET("/maintain", handler.QueryMaintainList)
		getNormalUserGroupManger.GET("/userlist", handler.QueryUserList) // userlist.
		// Beware that user can only modify their own res.
		getNormalUserGroupManger.GET("/user", handler.QuerySelfUserInfo)
		getNormalUserGroupManger.POST("/user", handler.ModifySelfUserInfo)

	}

	getRootUserGroupManger := SetRoot.Group("/root").Use(Cors(), middleware.RequireAuth, middleware.IsCanModifiyDataUser)
	{
		getRootUserGroupManger.POST("/res", handler.ResMangerModify)
		getRootUserGroupManger.POST("/res/add", handler.ResManagerAdd)
		getRootUserGroupManger.POST("/maintain", handler.ModifyMaintainList)
		getRootUserGroupManger.POST("/maintain/add", handler.AddMaintainList)
		getRootUserGroupManger.DELETE("/res", handler.ResmanagerDelete)
		getRootUserGroupManger.DELETE("/maintain", handler.DelMaintainList)
	}

	getHighestUserGroupManger := SetRoot.Group("/admin").Use(Cors(), middleware.RequireAuth, middleware.IsCanModifiyDataUser)
	{
		getHighestUserGroupManger.POST("/register", handler.Register)
		getHighestUserGroupManger.DELETE("/register", handler.DelUserList)
	}
	// REPORT HERE, GUEST CAN ALSO MODIFY DATA.
	SetRoot.POST("/root/report/add", handler.AddReport).Use(Cors(), middleware.RequireAuth) // set for no auth.
	SetRoot.POST("/root/report", handler.ModifyReport).Use(Cors(), middleware.RequireAuth)
	SetRoot.DELETE("/root/report", handler.DeleteReport).Use(Cors(), middleware.RequireAuth)

}
