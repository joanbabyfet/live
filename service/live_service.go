package service

import (
	"fmt"
	"live/config"
	"live/dto"
	"live/model"
	"net/http"
	"time"
)

type LiveService struct {
	Client *http.Client
}

// 构造函数
func NewLiveService() *LiveService {
	return &LiveService{
		Client: &http.Client{},
	}
}

func (s *LiveService) CreateRoom(req dto.CreateRoomReq) (*dto.LiveRoomResp, error) {

    roomID := fmt.Sprintf(
        "room_%d",
        time.Now().UnixNano(),
    )

    streamName := fmt.Sprintf(
        "live_%d",
        req.UID,
    )

    streamKey := fmt.Sprintf(
        "sk_%d_%d",
        req.UID,
        time.Now().Unix(),
    )

    room := model.LiveRoom{
        UID:     req.UID,
        RoomID:     roomID,
        Title:      req.Title,
        StreamName: streamName,
        StreamKey:  streamKey,
        Status:     "offline",
        CreateTime: time.Now().Unix(),
    }

    err := config.DB.Create(&room).Error

    if err != nil {
        return nil, err
    }

    pushURL := fmt.Sprintf(
        "rtmp://localhost/live/%s?key=%s",
        streamName,
        streamKey,
    )

    playURL := fmt.Sprintf(
        "http://localhost:8080/live/%s.m3u8",
        streamName,
    )

    return &dto.LiveRoomResp{
        RoomID:    roomID,
        PushURL:   pushURL,
        PlayURL:   playURL,
        StreamKey: streamKey,
    }, nil
}

//获取直播间列表
func (s *LiveService) GetLiveList() ([]model.LiveRoom, error) {

    var rooms []model.LiveRoom

    err := config.DB.
        Where("status = ?", "live").
        Find(&rooms).Error

    return rooms, err
}

//获取直播间详情
func (s *LiveService) GetLiveDetail(roomID string) (*model.LiveRoom, error) {

    var room model.LiveRoom

    err := config.DB.
        Where("room_id = ?", roomID).
        First(&room).Error

    return &room, err
}

// 根据 stream_name 获取直播间
func (s *LiveService) GetRoomByStream(streamName string) (*model.LiveRoom, error) {

    var room model.LiveRoom

    err := config.DB.
        Where("stream_name = ?", streamName).
        First(&room).Error

    return &room, err
}

//更新直播间状态
func (s *LiveService) UpdateLiveStatus(streamName string, status string) error {

    return config.DB.
        Model(&model.LiveRoom{}).
        Where("stream_name = ?", streamName).
        Update("status", status).Error
}