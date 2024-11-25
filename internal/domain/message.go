package domain

import "time"

type Message struct {
	Id        int64     `json:"id" bson:"_id,omitempty"`
	SenderId  int64     `json:"sender_id" bson:"sender_id"`
	RoomId    int64     `json:"room_id" bson:"room_id"`
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at bson:"updated_at,omitempty"`
}
