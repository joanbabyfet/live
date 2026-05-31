package router

import (
	"live/controller"
	"live/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(
    RoomController *controller.RoomController,
    SRSController *controller.SRSController,
) *gin.Engine {

    r := gin.Default()
    
    // 静态资源
    r.Static("/static", "./static")

    api := r.Group("/api")

    //模拟登录
    api.Use(middleware.Auth())
    {
        api.POST("/room/create", RoomController.Add)
        api.GET("/room/list", RoomController.Index)
        api.GET("/room/:id", RoomController.Detail)
        api.POST("/room/start", RoomController.StartLive)
        api.POST("/room/stop", RoomController.StopLive)
    }

    hook := r.Group("/hooks")
    {
        hook.POST("/on_publish", SRSController.OnPublish)
        hook.POST("/on_unpublish", SRSController.OnUnPublish)
    }

    return r
}