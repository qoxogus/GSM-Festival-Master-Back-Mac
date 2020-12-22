package database

import "time"

//User struct
type User struct {
	Classnum  string    `gorm:"primary_key; auto_increment:true" json:"classnum"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);noe null" json:"email"`
	Pw        string    `gorm:"type:varchar(255);not null" json:"pw"`
	IsManager bool      `gorm:"not null" sql:"DEFAULT:false" json:"is_manager"`
	JoinedAt  time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp" json:"joined_at"`
}
