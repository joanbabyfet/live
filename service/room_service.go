package service

import (
	"errors"
	"fmt"
	"live/config"
	"live/dto"
	"live/model"
	"net/http"
	"strconv"
	"time"
)

type RoomService struct {
	Client *http.Client
}

// 构造函数
func NewRoomService() *RoomService {
	return &RoomService{
		Client: &http.Client{},
	}
}

//创建
func (s *RoomService) Create(req dto.CreateRoomReq) (*dto.RoomResp, error) {

    roomID := fmt.Sprintf(
        "room_%d",
        time.Now().UnixNano(),
    )

    streamName := fmt.Sprintf(
        "live_%s",
        req.UID,
    )

    streamKey := fmt.Sprintf(
        "sk_%s_%d",
        req.UID,
        time.Now().Unix(),
    )

    room := model.Room{
        UID:        req.UID,
        RoomID:     roomID,
        Title:      req.Title,
        StreamName: streamName,
        StreamKey:  streamKey,
        Status:     0,
        CreateTime: time.Now().Unix(),
    }

    //写入数据库
    if err := config.DB.Create(&room).Error; err != nil {
		return nil, err
	}

    //写入缓存
    pipe := config.RDB.Pipeline()

	pipe.HSet(
		config.Ctx,
		"room:"+room.RoomID, //room:{roomID}
		"uid", room.UID,
		"room_id", room.RoomID,
		"title", room.Title,
		"stream_name", room.StreamName,
		"stream_key", room.StreamKey,
		"status", room.Status,
		"create_time", room.CreateTime,
	)

	// 在线房间列表索引
	pipe.SAdd(
		config.Ctx,
		"room:online",
		room.RoomID,
	)

	_, err := pipe.Exec(config.Ctx)

	if err != nil {
		return nil, err
	}

    //获取推流地址
    pushURL := s.GetPushURL(room)
    //获取播放地址
    playURL := s.GetPlayURL(room)

    return &dto.RoomResp{
        RoomID:    roomID,
        PushURL:   pushURL,
        PlayURL:   playURL,
        StreamKey: streamKey,
    }, nil
}

//获取直播间列表
func (s *RoomService) List() ([]model.Room, error) {

    var rooms []model.Room

    err := config.DB.
        Where("status = ?", "live").
        Find(&rooms).Error

    return rooms, err
}

//获取直播间详情
func (s *RoomService) GetByID(roomID string) (*model.Room, error) {

    // var room model.Room

    // err := config.DB.
    //     Where("room_id = ?", roomID).
    //     First(&room).Error

    // return &room, err

    data, err := config.RDB.
		HGetAll(config.Ctx, "room:"+roomID).
		Result()

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("room not found")
	}

    status, _ := strconv.Atoi(data["status"])
    createTime, _ := strconv.ParseInt(
        data["create_time"],
        10,
        64,
    )

	room := &model.Room{
		RoomID:     data["room_id"],
        UID:        data["uid"],
        Title:      data["title"],
        StreamName: data["stream_name"],
        StreamKey:  data["stream_key"],
        Status:     int8(status),
        CreateTime: createTime,
	}

	return room, nil
}

// 根据 stream_name 获取直播间
func (s *RoomService) GetByStreamName(streamName string) (*model.Room, error) {

    var room model.Room

    err := config.DB.
        Where("stream_name = ?", streamName).
        First(&room).Error

    return &room, err
}

//更新直播间状态
func (s *RoomService) UpdateStatus(streamName string, status int8) error {

    return config.DB.
        Model(&model.Room{}).
        Where("stream_name = ?", streamName).
        Update("status", status).Error
}

//开始直播
func (s *RoomService) StartLive(roomID string) (error) {
    pipe := config.RDB.Pipeline()

	pipe.HSet(
		config.Ctx,
		"room:"+roomID,
		"status",
		1,
	)
    
    //加入推流状态 room:stream:{roomID}
    pipe.Set(config.Ctx, "room:stream:"+roomID, 1, 0)

     //加入在线房间
	pipe.SAdd(
		config.Ctx,
		"room:online",
		roomID,
	)

	_, err := pipe.Exec(config.Ctx)

	return err
}

//停止直播
func (s *RoomService) StopLive(roomID string) (error) {
    fmt.Println("roomID =", roomID)

    pipe := config.RDB.Pipeline()

	pipe.HSet(
		config.Ctx,
		"room:"+roomID,
		"status",
		0,
	)

    //移除推流状态 room:stream:{roomID}
    pipe.Del(config.Ctx,  "room:stream:"+roomID)

    //移除在线房间
	pipe.SRem(
		config.Ctx,
		"room:online",
		roomID,
	)

	_, err := pipe.Exec(config.Ctx)

	return err
}

// 在线房间列表
func (s *RoomService) OnlineRoom() []string {
    val, _ := config.RDB.SMembers(config.Ctx, "room:online").Result()

    return val
}

// 设置观众数
func (s *RoomService) SetViewers(roomID string, count int) {
    config.RDB.Set(config.Ctx, "room:viewers:"+roomID, count, 0)
}

// 获取观众数
func (s *RoomService) GetViewers(roomID string) int {
    val, err := config.RDB.Get(config.Ctx, "room:viewers:"+roomID).Int()

    if err != nil {
        return 0
    }
    return val
}

//获取推流地址(主播用)
func (s *RoomService) GetPushURL(room model.Room) (string) {
	return fmt.Sprintf(
        "%s/%s?key=%s",
        config.App.PushURL,
        room.StreamName,
		room.StreamKey,
    ) 
}

//获取播放地址(观众用)
func (s *RoomService) GetPlayURL(room model.Room) (string) {
	return fmt.Sprintf(
        "%s/%s.flv",
        config.App.PlayURL,
        room.StreamName,
    ) 
}