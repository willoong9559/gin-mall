package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/willoong9559/gin-mall/api/v1"
	"github.com/willoong9559/gin-mall/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户相关
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 商品相关
		v1.GET("carousels", api.ListCarousels) //轮播图
	}

	authed := v1.Group("/") //需要登陆保护
	authed.Use(middleware.JWT())
	{
		// 用户操作
		authed.PUT("user", api.UserUpdate)
		authed.POST("avatar", api.UploadAvatar)
		authed.POST("user/sending-email", api.SendEmail)
		authed.POST("user/valid-email", api.ValidEmail)

		// 显示金额
		authed.POST("money", api.ShowMoney)
	}
	return r
}