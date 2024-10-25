package chatdomain

type Room struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique; not null"`
}

type RoomCreate struct {
	Name string `json:"name"`
}
