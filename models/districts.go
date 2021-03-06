package models

import "time"

type Districts struct {
	Id           int     `json:"id" validate:"required"`
	CreatedBy    string     `json:"created_by" validate:"required"`
	CreatedDate  time.Time  `json:"created_date" validate:"required"`
	ModifiedBy   *string    `json:"modified_by"`
	ModifiedDate *time.Time `json:"modified_date"`
	DeletedBy    *string    `json:"deleted_by"`
	DeletedDate  *time.Time `json:"deleted_date"`
	IsDeleted    int        `json:"is_deleted" validate:"required"`
	IsActive     int        `json:"is_active" validate:"required"`
	DistrictsName  string `json:"districts_name"`
	CityId *int `json:"city_id"`
}
type NewCommandDistricts struct {
	Id          int `json:"id" validate:"required"`
	DistrictsName  string `json:"districts_name"`
	CityId *int `json:"city_id"`
}

type DistrictsDto struct {
	Id          int `json:"id" validate:"required"`
	DistrictsName  string `json:"districts_name"`
	CityId *int `json:"city_id"`
}
type DistrictsWithPagination struct {
	Data []*DistrictsDto  `json:"data"`
	Meta *MetaPagination `json:"meta"`
}