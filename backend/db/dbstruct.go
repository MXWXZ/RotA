package db

// User database map
type User struct {
	ID       int32  `gorm:"auto_increment;primary_key"`
	UserName string `gorm:"size:32;not null"`
	UserPass string `gorm:"size:64;not null"`
}
