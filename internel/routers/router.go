package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/willoong9559/gin-mall/docs"
	api "github.com/willoong9559/gin-mall/internel/api/v1"
	"github.com/willoong9559/gin-mall/internel/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Session("some_key"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		//  用户相关
		v1.GET("/captcha", api.GetCaptcha)
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		// 商品相关
		v1.GET("/carousels", api.ListCarousels) //轮播图
		v1.GET("/listProducts", api.ListProducts)
		v1.GET("/searchProducts", api.SearchProducts)
	}

	authUser := v1.Group("/") //需要登陆保护
	authUser.Use(middleware.JWT())
	{
		// 用户操作
		authUser.PUT("/user", api.UserUpdate)
		authUser.POST("/avatar", api.UploadAvatar)
		authUser.POST("/user/sending-email", api.SendEmail)
		authUser.POST("/user/valid-email", api.ValidEmail)

		// 显示金额
		authUser.POST("/money", api.ShowMoney)

		// 商品操作
		authUser.POST("/product", api.CreateProduct)
	}
	return r
}
