package route

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goAdminStudy/controller"
	"goAdminStudy/middlewires"
	"goAdminStudy/utils"
	"net/http"
)

func InitRoute() *gin.Engine {

	engine := gin.New()

	engine.Use(middlewires.LoggerMid())
	engine.Use(middlewires.GlobalError)
	engine.Use(middlewires.NoCache)
	engine.Use(middlewires.Secure)

	//
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, utils.SuccessResult("pong"))
	})
	engine.POST("/user/login")
	engine.POST("/user/register")

	// 需要登陆验证路由组
	authGroup := engine.Group("/adm", middlewires.JwtMiddle())

	taskGroup := authGroup.Group("/task")
	{
		taskGroup.GET("/get")
		taskGroup.GET("/find")
		taskGroup.POST("/save")
		taskGroup.DELETE("/delete")
		taskGroup.PUT("/update")
	}

	kafkaGroup := authGroup.Group("/kafka")
	{
		kafkaController := controller.KafkaController{}
		kafkaGroup.POST("/save", kafkaController.Add)
		kafkaGroup.GET("/list", kafkaController.List)
	}

	//swager
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return engine

}
