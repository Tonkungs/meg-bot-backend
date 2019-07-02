package models

// import  "go.mongodb.org/mongo-driver/bson"
const (
	// CollectionUser  Messages collection
	CollectionMessages = "Messages"
)

// Messages ข้อความ
type Messages struct {
	UserID     string `json:"user_id"`
	UserLineID string `json:"user_line_id"`
	Type       string `json:"type"`
	Text       string `json:"text"`
	CreatedAt  int64  `json:"created_at"`
	UpdateAt   int64  `json:"update_at"`
}
