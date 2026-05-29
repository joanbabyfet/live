package main

import (
	"live/config"
	"live/controller"
	"live/router"
	"live/service"
	"log"
)

func main() {

    // 初始化配置
    config.InitConfig()

    // 初始化数据库
    config.InitDB()

    // 创建 service 实例
    liveService := service.NewLiveService()

    // 创建 controller 实例
    liveController := controller.NewLiveController(
        liveService,
    )
    SRSController := controller.NewSRSController(
        liveService,
    )
    // 路由初始化
    r := router.InitRouter(
        liveController,
        SRSController,
    )

    log.Println("服务启动:", config.App.Port)

    err := r.Run(":" + config.App.Port)

	if err != nil {
		panic(err)
	}
}