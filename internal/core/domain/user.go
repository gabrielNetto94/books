package domain

type User struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
