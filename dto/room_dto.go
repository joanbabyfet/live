package dto

type StartLiveReq struct {
	RoomID string `json:"room_id"`
}

type StopLiveReq struct {
	RoomID string `json:"room_id"`
}

type CreateRoomReq struct {
    UID     string `json:"-"` //表示前端不能传 UID
    Title   string `json:"title" binding:"required"`
}

type RoomResp struct {
    RoomID    string `json:"room_id"`
    PushURL   string `json:"push_url"`
    PlayURL   string `json:"play_url"`
    StreamKey string `json:"stream_key"`
}