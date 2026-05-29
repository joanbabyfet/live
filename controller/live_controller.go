package controller

import (
	"live/common"
	"live/dto"
	"live/service"

	"github.com/gin-gonic/gin"
)

type LiveController struct {
	LiveService *service.LiveService
}

//构造函数
func NewLiveController(s *service.LiveService) *LiveController {
	return &LiveController{
		LiveService: s,
	}
}

//创建直播间
func (c *LiveController) CreateRoom(ctx *gin.Context) {

    var req dto.CreateRoomReq

    if err := ctx.ShouldBindJSON(&req); err != nil {

        common.Fail(ctx, -1, err.Error(), nil)

        return
    }

    resp, err := c.LiveService.CreateRoom(req)

    if err != nil {

        common.Fail(ctx, -2, err.Error(), nil)

        return
    }

    common.Success(ctx, gin.H{"data": resp})
}

//创建直播间列表
func (c *LiveController) LiveList(ctx *gin.Context) {

    list, err := c.LiveService.GetLiveList()

    if err != nil {

        common.Fail(ctx, -1, err.Error(), nil)

        return
    }

    common.Success(ctx, gin.H{"list": list})
}

//创建直播间详情
func (c *LiveController) LiveDetail(ctx *gin.Context) {

    roomID := ctx.Param("room_id")

    data, err := c.LiveService.GetLiveDetail(roomID)

    if err != nil {

        common.Fail(ctx, -1, err.Error(), nil)

        return
    }

    common.Success(ctx, gin.H{"data": data})
}