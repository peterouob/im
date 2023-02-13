package router

import (
	"gin-chat/docs"
	"gin-chat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//靜態資源
	r.Static("/assets", "assets/")
	r.LoadHTMLGlob("views/**/*.html")
	//首頁
	r.GET("/", service.GetIndex)
	//用戶模塊
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)

	//發送消息
	r.GET("/user/sendMessage", service.SendMessage)
	r.GET("/user/sendUserMessage", service.SendUserMessage) //死鎖問題！
	return r
}
