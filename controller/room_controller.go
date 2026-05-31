package controller

import (
	"live/common"
	"live/dto"
	"live/service"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	RoomService *service.RoomService
}

//构造函数
func NewRoomController(s *service.RoomService) *RoomController {
	return &RoomController{
		RoomService: s,
	}
}

//创建直播间, ctx为请求容器
func (c *RoomController) Add(ctx *gin.Context) {

    var req dto.CreateRoomReq

    if err := ctx.ShouldBindJSON(&req); err != nil {
        common.Fail(ctx, -1, err.Error(), nil)

        return
    }

    //从 Gin 的 Context（上下文）获取 userID, 上下文（Context）就是一次 HTTP 请求的“共享数据容器”
    req.UID = ctx.MustGet("userID").(string)
    
    resp, err := c.RoomService.Create(req)

    if err != nil {

        common.Fail(ctx, -2, err.Error(), nil)

        return
    }

    common.Success(ctx, resp)
}

//获取直播间列表
func (c *RoomController) Index(ctx *gin.Context) {
    list, err := c.RoomService.List()

    if err != nil {
        common.Fail(ctx, -1, err.Error(), nil)
        return
    }
    common.Success(ctx, gin.H{"list": list})
}

//获取直播间详情
func (c *RoomController) Detail(ctx *gin.Context) {
    id := ctx.Param("id")

    data, err := c.RoomService.GetByID(id)

    if err != nil {
        common.Fail(ctx, -1, err.Error(), nil)
        return
    }
    common.Success(ctx, data)
}

//开始直播
func (c *RoomController) StartLive(ctx *gin.Context) {
     var req dto.StartLiveReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.Fail(ctx, -1, err.Error(), nil)
		return
	}

    err := c.RoomService.StartLive(req.RoomID)

    if err != nil {
        common.Fail(ctx, -1, err.Error(), nil)
        return
    }
    common.Success(ctx, gin.H{"status": 1})
}

//结束直播
func (c *RoomController) StopLive(ctx *gin.Context) {
    var req dto.StopLiveReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.Fail(ctx, -1, err.Error(), nil)
		return
	}

    err := c.RoomService.StopLive(req.RoomID)

    if err != nil {
        common.Fail(ctx, -1, err.Error(), nil)
        return
    }
    common.Success(ctx, gin.H{"status": 0})
}