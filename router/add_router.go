package router

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/controllers"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	"path"
)

//addRouter add routers
func addRouter(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var adapter *gormadapter.Adapter
	db := common.DatabaseInfo
	var connectInfo string
	connectInfo = db.Connect
	wr, _ := os.Getwd()
	adapter = gormadapter.NewAdapter(db.DBType, connectInfo, true)
	common.Enforcer = casbin.NewEnforcer(path.Join(wr, "conf", "casbin_rbac_model.conf"), adapter)
	common.Enforcer.LoadPolicy()
	authRouter := router.Group("/auth")
	//authRouter.Use(middles.CasbinJwtAuthorize(common.Enforcer))
	{
		// router no need login ,authRouter need user login
		router.GET("/", controllers.IndexGet)
		router.POST("/login", controllers.Login)
		router.POST("/register", controllers.UserRegister)
		router.POST("/login/wechat", controllers.LoginWechat)
		router.GET("/third/wechat", controllers.GetOpenId)

		//微信  pays
		authRouter.POST("/pay/wxapp", controllers.WxAppPay)
		authRouter.POST("/pay/asyncback", controllers.WxAppPayAsyncBack)
		authRouter.POST("/weixin_upload", controllers.WeiXinUpload)

		// refresh token
		authRouter.GET("/refresh", controllers.RefreshToken)

		// user
		authRouter.GET("/self", controllers.SelfInfo)
		authRouter.PUT("/self", controllers.SelfInfoUpdate)
		authRouter.GET("/user/:id", controllers.UserGet)
		authRouter.PUT("/user/:id", controllers.UserUpdate)
		authRouter.GET("/users", controllers.UserGetAll)
		authRouter.POST("/logout", controllers.Logout)
		authRouter.POST("/user/register", controllers.UserRegister)

		// menu
		authRouter.GET("/menu/:id", controllers.MenuGet)
		authRouter.PUT("/menu/:id", controllers.MenuUpdate)
		authRouter.GET("/menutree/", controllers.MenuTree)
		authRouter.GET("/usermenu", controllers.UserMenu)
		authRouter.GET("/usermenutree/:user_id", controllers.UserMenuTree)

		// role
		authRouter.POST("/role", controllers.RoleCreate)
		authRouter.GET("/role/:id", controllers.RoleGet)
		authRouter.PUT("/role/:id", controllers.RoleUpdate)
		authRouter.DELETE("/role/:id", controllers.RoleDelete)
		authRouter.GET("/roles", controllers.RoleAll)

		// 公司管理

	}
}
