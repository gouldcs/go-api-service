package models

type Player struct {
	PlayerID uint `gorm:"primaryKey"`
	Firstname string `gorm:"column:first_name"`
	Lastname string `gorm:"column:last_name"`
	JerseyNumber string
	Team Team `gorm:"foreignKey:TeamID"`
}