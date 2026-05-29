package model

type LiveRoom struct {
    ID         int64     `gorm:"primaryKey" json:"id"`
    UID     int64     `json:"uid"`
    RoomID     string    `json:"room_id"`
    Title      string    `json:"title"`
    StreamName string    `json:"stream_name"`
    StreamKey  string    `json:"stream_key"`
    Status     string    `json:"status"`
    CreateTime  int64 `json:"create_time"`
}