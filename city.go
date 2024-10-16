package main

type City struct {
	CityID string `json:"cityId"`
	CityName string `json:"cityName"`
	CityAlias string `json:"cityAlias"`
	CityState string `json:"cityState"`
}

var cities = []City {
	{CityID: "1", CityName: "Seattle", CityAlias: "SEA", CityState: "WA"},
	{CityID: "2", CityName: "Los Angeles", CityAlias: "LA", CityState: "CA"},
	{CityID: "3", CityName: "Houston", CityAlias: "HOU", CityState: "TX"},
	{CityID: "4", CityName: "New York", CityAlias: "NY", CityState: "NY"},
	{CityID: "5", CityName: "Miami", CityAlias: "MIA", CityState: "FL"},
}

func getCities() []City {
	return cities
}