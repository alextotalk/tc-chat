package domain

type UserRoom struct {
	Id       int64 `json:"id" bson:"_id,omitempty,pk"`
	UserId   int64 `json:"user_id" bson:"user_id, omitempty,pk"`
	RoomId   int64 `json:"sender_id" bson:"sender_id, omitempty,pk"`
	IsBanned bool  `json:"is_banned" bson:"is_banned,omitempty"`
}
