package models

import "time"

type HurunguReponse struct {
	List     []HunguData `json:"list"`
	ListSize int         `json:"listSize"`
}
type HunguData struct {
	AssetCode    string    `json:"assetCode"`
	AssetName    string    `json:"assetName"`
	MonthsUsed   int       `json:"monthsUsed"`
	CommenceDate time.Time `json:"commenceDate"`
	Qty          float32   `json:"qty"`
	Total        float32   `json:"total"`
	Price        float32   `json:"price"`
	GroupCode    string    `json:"groupCode"`
	GroupName    string    `json:"groupName"`
	UnitCode     string    `json:"unitCode"`
	UnitName     string    `json:"unitName"`
	OwnerCode    string    `json:"ownerCode"`
	OwnerName    string    `json:"ownerName"`
	CategoryCode string    `json:"categoryCode"`
	CategoryName string    `json:"categoryName"`
	LocationCode string    `json:"locationCode"`
	LocationName string    `json:"locationName"`
}
