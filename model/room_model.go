package model

//直播间
type Room struct {
    ID          int64       `gorm:"primaryKey;description(ID)" json:"id"`
    RoomID      string      `gorm:"primaryKey;description(房间标识)" json:"room_id"`
    UID         string      `gorm:"description(主播ID)" json:"uid"`
    Title       string      `gorm:"description(名称)" json:"title"`
    StreamName  string      `gorm:"description(流名)" json:"stream_name"`
    StreamKey   string      `gorm:"description(密钥)" json:"stream_key"`
    Status      int8        `gorm:"description(状态 0:离线 1:在线)" json:"status"`
    CreateTime  int64       `gorm:"description(创建时间)" json:"create_time"`
}

//定义表名
func (m *Room) TableName() string {
	return "room"
}