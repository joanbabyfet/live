package dto

type SRSHookReq struct {
    Action   string `json:"action"` //比如 on_publish
    ClientID string `json:"client_id"` //客户端id 比如 ClientID:s136w5u3
    IP       string `json:"ip"`
    Vhost    string `json:"vhost"`
    App      string `json:"app"` //比如 live
    Stream   string `json:"stream"` //比如 live_1003
    Param    string `json:"param"` //比如 ?key=sk_1003_1779876646
}