package dto

type CreateRoomReq struct {
    UID int64  `json:"uid" binding:"required"`
    Title  string `json:"title" binding:"required"`
}

type LiveRoomResp struct {
    RoomID    string `json:"room_id"`
    PushURL   string `json:"push_url"`
    PlayURL   string `json:"play_url"`
    StreamKey string `json:"stream_key"`
}