package models

type Session struct {
	Token        string `json:"token"`
	UserId       int    `json:"user_id"`
	IpAddress    string `json:"ip_address"`
	UserAgent    string `json:"user_agent"`
	LastActivity int64  `json:"last_activity"`
}
