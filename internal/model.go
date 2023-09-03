package internal

type proxyRequest struct {
	Payload      string   `json:"payload" required:"true"`
	UserId       string   `json:"user_id" required:"true"`
	Destinations []string `json:"destinations" required:"true"`
}
