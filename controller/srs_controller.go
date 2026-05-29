package controller

import (
	"live/dto"
	"live/service"
	"log"

	"github.com/gin-gonic/gin"
)

type SRSController struct {
	LiveService *service.LiveService
}

//构造函数
func NewSRSController(s *service.LiveService) *SRSController {
	return &SRSController{
		LiveService: s,
	}
}

//直播开始(主播点击obs开始直播), SRS收到推流后调用, 调用 POST /hooks/on_publish
func (c *SRSController) OnPublish(ctx *gin.Context) {

    var req dto.SRSHookReq

    if err := ctx.ShouldBindJSON(&req); err != nil {

        ctx.JSON(400, gin.H{
            "code": 1,
        })

        return
    }

    // 打印完整 req struct 字段名与值
    log.Printf("OnPublish req: %+v\n", req)

    // 推流鉴权
    room, err := c.LiveService.GetRoomByStream(req.Stream)
    if err != nil {
        ctx.JSON(404, gin.H{
            "code": 404,
            "msg": "直播间不存在",
        })
        return
    }

    //验证 key
    if req.Param != "?key="+room.StreamKey {
        log.Println("推流鉴权失败:", req.Stream)

        ctx.JSON(403, gin.H{
            "code": 403,
            "msg": "invalid stream key",
        })
        return
    }

    log.Println("直播开始:", req.Stream)

    c.LiveService.UpdateLiveStatus(req.Stream, "live")

    ctx.JSON(200, gin.H{
        "code": 0,
    })
}

//直播结束(主播点击obs结束直播)
func (c *SRSController) OnUnPublish(ctx *gin.Context) {

    var req dto.SRSHookReq

    if err := ctx.ShouldBindJSON(&req); err != nil {

        ctx.JSON(400, gin.H{
            "code": 1,
        })

        return
    }

    // 打印完整 req struct 字段名与值
    log.Printf("OnUnPublish req: %+v\n", req)

    log.Println("直播结束:", req.Stream)

    c.LiveService.UpdateLiveStatus(req.Stream, "offline")

    ctx.JSON(200, gin.H{
        "code": 0,
    })
}