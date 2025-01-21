package models

type City struct {
	CityID uint `gorm:"primaryKey;autoIncrement" json:"city_id"`
	CityName string `gorm:"column:city_name" json:"city_name"`
	CityAlias string `gorm:"column:city_alias" json:"city_alias"`
	CityState string `gorm:"column:city_state" json:"city_state"`
}