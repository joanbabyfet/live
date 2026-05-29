package router

import (
	"live/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(
    liveController *controller.LiveController,
    SRSController *controller.SRSController,
) *gin.Engine {

    r := gin.Default()
    
    // 静态资源
    r.Static("/static", "./static")

    live := r.Group("/live")
    {
        live.POST("/create", liveController.CreateRoom)
        live.GET("/list", liveController.LiveList)
        live.GET("/:room_id", liveController.LiveDetail)
    }

    hook := r.Group("/hooks")
    {
        hook.POST("/on_publish", SRSController.OnPublish)
        hook.POST("/on_unpublish", SRSController.OnUnPublish)
    }

    return r
}