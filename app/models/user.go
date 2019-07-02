package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionUser  users collection
	CollectionUser = "users"
)

// type Usered interface {
// 	getDisplay() string
// }

// User User
type User struct {
	// ID           primitive.ObjectID `json:"_id", bson:"_id,omitempty"`
	ID           primitive.ObjectID `bson:"_id"`
	UserID       string             `json:"user_id"`
	DisplayPhoto string             `json:"display_photo"`
	DisplayName  string             `json:"display_name"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdateAt     time.Time          `json:"update_at"`
	// usered		 Usered
}

// GetDisplay nested in a struct
func (u User) GetDisplay() string {
	return u.DisplayPhoto + "-" + u.DisplayName
}

// Create For use
// func (user *User) Create() string {
// 	return "123456"
// }
