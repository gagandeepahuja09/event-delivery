package internal

type proxyRequest struct {
	Payload      string   `json:"payload"`
	UserId       string   `json:"user_id"`
	Destinations []string `json:"destinations"`
}
