package router

import (
	"github.com/gin-gonic/gin"
	"younghe/controller/account"
	"younghe/middleware"
)
import "younghe/lib/corex"
import "younghe/controller/demo"

func InitRouter() *gin.Engine {
	//gin.SetMode("release")
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(corex.Handle(middleware.Recovery()))
	engine.Use(corex.Handle(middleware.RequestId()))

	// 路由
	engine.GET("/ping", corex.Handle(demo.Ping))
	engine.GET("/api/mysql/getAll", corex.Handle(demo.GetAll))
	engine.GET("/api/mysql/getPage", corex.Handle(demo.PageList))
	engine.POST("/api/mysql/createOne", corex.Handle(demo.CreateOne))
	engine.GET("/api/mysql/getOne", corex.Handle(demo.GetOne))
	engine.POST("/api/mysql/deleteOne", corex.Handle(demo.DeleteOne))
	engine.GET("/api/redis/get", corex.Handle(demo.RedisGet))
	engine.POST("/api/redis/set", corex.Handle(demo.RedisSet))
	engine.GET("/api/panic", corex.Handle(demo.PanicTest))

	engine.POST("/api/account/login", corex.Handle(account.Login))
	engine.POST("/api/account/register", corex.Handle(account.Register))
	accountNeedLoginGroup := engine.Group("/api/account", corex.Handle(middleware.LoginAuth()))
	{
		accountNeedLoginGroup.GET("/getOwn", corex.Handle(account.GetOwn))
		accountNeedLoginGroup.POST("/logout", corex.Handle(account.Logout))
	}

	return engine
}
