package router

import (
	"GraduationDesign/controller"
	"GraduationDesign/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解决跨域问题的中间件
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
func InitRouter() {
	r := gin.Default()
	r.Use(Cors())
	r.Static("/file", "./file")
	r.Static("/pictures", "./pictures")
	r.Static("/photos", "./photos")
	GetRouters(r)
	r.Run(":80")
}

func GetRouters(r *gin.Engine) {
	//游客总模块
	common := r.Group("/common")
	{
		//获取公告		表单传入参数：无 返回值：公告
		common.GET("/notice", controller.GetNotice)
		//注册总模块
		register := common.Group("/register")
		{
			//注册		表单传入参数：1.用户名username 2.密码password 3.电话号码phonenumber 4.验证码code 返回值：无
			register.POST("", controller.Register)
			//邮箱注册
			//register.POST("/email", controller.Register2)
			//获取验证码	路径参数传入：1.电话号码phonenumber 返回值：无
			register.GET("/:phonenumber", controller.GetVerification)
			//获取验证码2 路径参数传入：1.邮箱email 返回值：无
			//register.GET("/email/:email", controller.GetVerification2)
		}

		//登录总模块
		login := common.Group("/login")
		{
			//获取验证码		路径参数传入：1.电话号码phonenumber 返回值：无
			login.GET("/:phonenumber", controller.GetVerification)
			//密码登录 		表单传入参数：1.用户名username 2.密码password 返回值：用户信息
			login.POST("", controller.Login)
			//电话验证码登录		表单传入参数：1.电话号码phonenumber 2.验证码code 返回值：用户信息
			login.POST("/p", controller.LoginByPhoneNumber)
		}
	}

	//用户总模块
	user := r.Group("/user", middleware.LoginNeed)
	{
		//用户信息模块
		userInfo := user.Group("/userInfo")
		{
			//获得用户信息	表单传入参数：1.无 返回值:用户信息
			userInfo.GET("/", controller.GetUserInfo)
			//修改用户名		表单传入参数：1.新用户名username 返回值:新用户名
			userInfo.PUT("/username", controller.UpdateUsername)
			//修改头像		表单传入参数：1.头像参数avatarNumber 返回值:头像地址
			userInfo.PUT("/avatar", controller.UpdateAvatar)
			//修改用户个人简介	表单传入参参数：1.自我介绍内容 返回值:自我介绍内容
			userInfo.PUT("/introduction", controller.UpdateIntroduction)
			//修改用户信息
			userInfo.PUT("", controller.UpdateUserInfo)
		}
		//商品模块
		goods := user.Group("/goods")
		{
			//获得所有商品信息 需要返回的数据：1.商品名 2.商品所属人员 3.商品图片（一张） 4.价格
			goods.GET("/goodsdata", controller.GetGoodsData)
			//获得所有商品信息 需要返回的数据：1.商品名 2.商品所属人员 3.商品图片（一张） 4.价格 传入参数页数
			goods.POST("/goodsdata", controller.GetGoodsDataByPage)
			//获得所有商品信息 需要返回的数据：1.商品名 2.商品所属人员 3.商品图片（一张） 4.价格 传入参数页数
			goods.POST("/goodsdata/category", controller.GetGoodsDataByCategory)
			//获得商品信息	路径传入参数：1.商品id(锁定)  返回值：商品信息
			goods.GET("/:gid", controller.GetGoodsInfo)
			//添加商品		表单传入参数：1.商品名称(用户填写) 2.分类id（用户填写） 3.用户id(锁定)
			//4.图片地址(通过上传图片后的返回值获得) 5.商品简介(用户填写) 6.使用时间（用户填写） 7.价格（用户填写）  返回值：商品信息
			goods.POST("", controller.UploadGoods)
			//添加商品图片	表单传入参数：1.商品id(锁定) 2.图片二进制数据(用户上传)
			goods.POST("/picture", controller.UploadPictures)
			//获得某一用户的商品信息 需要展示的数据：1.商品名 2.商品所属人员 3.商品图片（一张） 4.价格
			goods.POST("/goodsdata/:uid", controller.GetGoodsByUid)
		}
		//购物车模块
		cart := user.Group("/cart")
		{
			//获取用户购物车
			cart.GET("", controller.GetUserCart)
			//添加图片到购物车
			cart.POST("/:gid", controller.AddGoodInCart)
			//将物品移除购物车
			cart.DELETE("/:gid", controller.RemoveGoodInCart)
		}
		//物品类别模块
		category := user.Group("/category")
		{
			//获得所有商品种类
			category.GET("", controller.GetAllCategory)
		}
		//投诉建议模块
		advice := user.Group("/advice")
		{
			//发起投诉建议
			advice.POST("", controller.AddAdvice)
			//获得所有建议种类
			advice.GET("/category", controller.GetAllAdviceCategory)
			//添加图片到商品内
			advice.POST("/photo", controller.UploadPhotos)
		}
		//浏览历史模块
		history := user.Group("history")
		{
			history.GET("", controller.GetHistory)
		}
		//收藏模块
		collection := user.Group("/collection")
		{
			collection.GET("", controller.GetCollection)
			//收藏/取消收藏 传入数据:1.商品id
			collection.POST("", controller.Collection)
		}
		//订单模块
		order := user.Group("/order")
		{
			order.GET("", controller.GetOrders)

			order.POST("", controller.CreateOrder)

			order.GET("/sell", controller.GetMySell)

			order.GET("/detail/:oid", controller.GetOrderDetail)

			order.PUT("", controller.UpdateOrder)
		}
		//认证模块
		indentification := user.Group("/identification")
		{
			//获取认证信息
			indentification.GET("", controller.GetIdentification)
			//新建认证信息
			indentification.POST("", controller.AddIdentification)
			//更新认证信息
			indentification.PUT("", controller.UpdateIdentification)
		}
	}
	//管理员模块
	admin := r.Group("/admin", middleware.LoginNeed, middleware.AdminNeed)
	{
		//公告模块
		notice := admin.Group("/notice")
		{
			//修改公告		表单传入参数：1.之前公告时间 2.公告更新的内容 返回值:公告内容
			notice.PUT("/", controller.UpdateNotice)
		}
		//用户信息模块
		userInfo := admin.Group("/userInfo")
		{
			//获得所有用户信息	返回值:所有用户信息
			userInfo.GET("/", controller.GetAllUserInfo)
		}
		goods := admin.Group("/goods")
		{
			//上架物品
			goods.POST("/allow/:gid", controller.AllowGoods)
			//下架物品
			goods.POST("/ban/:gid", controller.BanGoods)
			//获得所有物品信息
			goods.POST("", controller.GetAllGoodsData)
		}
		//商品种类模块
		category := admin.Group("/category")
		{
			//添加/修改商品种类 	传入参数：1.种类id,种类名称
			category.POST("", controller.AddCategory)
			//删除商品种类
			category.POST("/delete", controller.DeleteCategory)
		}
		//投诉建议模块
		advice := admin.Group("/advice")
		{
			//获取所有投诉建议
			advice.GET("", controller.GetAllAdvice)
			//根据建议id获取投诉建议
			advice.GET("/detail/:aid", controller.GetAdviceByAid)
			//确认阅读
			advice.PUT("/:aid", controller.ConfirmAdvice)
			//情况已阅读投诉建议
			advice.DELETE("", controller.DeleteAdvice)
		}
		//数据分析模块
		data := admin.Group("/data")
		{
			data.GET("", controller.GetData)
			data.GET("/user", controller.GetUserData)
			data.PUT("/user", controller.UpdateUserData)
			data.DELETE("/user/:uid", controller.DeleteUser)
		}
		identification := admin.Group("/identification")
		{
			//获取所有认证信息
			identification.POST("", controller.GetAllIdentification)
			//通过认证
			identification.PUT("/accept/:id", controller.AcceptIdentification)
			//拒绝认证
			identification.PUT("/refuse/:id", controller.RefuseIdentification)
		}
	}

}
