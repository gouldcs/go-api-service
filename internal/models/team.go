package models

type Team struct {
    TeamID uint `gorm:"primaryKey" json:"team_id"`
    CityID uint `gorm:"column:city_id" json:"city_id"`
    TeamName string `gorm:"column:team_name" json:"team_name"`
}