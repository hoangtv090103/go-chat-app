package authdomain

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
}

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
