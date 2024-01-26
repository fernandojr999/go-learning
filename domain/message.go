package domain

type Message struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Message  string `json:"message"`
	To       int    `json:"to_id"`
}
