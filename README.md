# Live Streaming System

Gin + SRS + OBS 多用户直播系统

---

# 技术栈

* Go + Gin
* MySQL
* SRS 5
* OBS Studio
* HTTP-FLV
* HLS
* Docker Compose

---

# 项目架构

```text
OBS
 ↓ RTMP推流
SRS
 ↓ Hook
Gin API
 ↓
MySQL

播放器:
HTTP-FLV / HLS
```

---

# 功能

* 多直播间
* OBS 推流
* HTTP-FLV 播放
* HLS 播放
* SRS Hook 回调
* 自动更新直播状态
* 直播列表
* 直播详情
* Stream 鉴权
* 静态播放器页面

---

# 项目目录

```text
project
├── controller
├── dto
├── model
├── router
├── service
├── static
│   └── player
│       ├── index.html
│       ├── player.js
│       └── flv-player.js
├── config
├── main.go
├── compose.yaml
├── srs.conf
└── README.md
```

---

# 环境变量

`.env`

```env
PORT=8081
```

---

# 启动 Gin

```bash
go run main.go
```

---

# 启动 SRS 与 mysql

```bash
docker compose up -d
```

---

# OBS 推流配置

服务器：

```text
rtmp://localhost/live
```

推流码：

```text
live_1003
```

完整 RTMP：

```text
rtmp://localhost/live/live_1003
```

---

# HTTP-FLV 播放地址

```text
http://127.0.0.1:8080/live/live_1003.flv
```

---

# HLS 播放地址

```text
http://127.0.0.1:8080/live/live_1003.m3u8
```

---

# 播放器页面

```text
http://localhost:8081/static/player/index.html
```

---

# Hook 回调

SRS 推流开始：

```text
POST /hooks/on_publish
```

SRS 推流结束：

```text
POST /hooks/on_unpublish
```

---

# Live API

## 创建直播间

```text
POST /live/create
```

---

## 直播列表

```text
GET /live/list
```

---

## 直播详情

```text
GET /live/:room_id
```

---

# Stream 鉴权

推流时：

```text
rtmp://localhost/live/live_1003?key=sk_xxx
```

Gin Hook 验证：

* stream_name
* stream_key

---

# HTTP-FLV

特点：

* 低延迟
* 1~3秒
* Chrome 支持 flv.js

适合：

* 真人直播
* 低延迟场景

---

# HLS

特点：

* 高兼容
* CDN友好
* 延迟较高

适合：

* 回放
* 手机端

---

# 后续扩展

* WebRTC
* 聊天室
* 在线人数
* 录播
* CDN
* 多节点
* JWT
* 后台管理系统

---

# License

[MIT License](https://github.com/joanbabyfet/live/blob/main/LICENSE)