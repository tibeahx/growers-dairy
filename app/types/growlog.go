package types

import "time"

type GrowLog struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	StrainID    int       `json:"strain_id"`
}

type CreateGrowLogReq struct {
	Name        string    `json:"name"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	StrainID    int       `json:"strain_id"`
}

type CreateGrowLogResp struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	StrainID    int       `json:"strain_id"`
}

type UpdateGrowLogReq struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	StrainID    int       `json:"strain_id"`
}

type UpdateGrowLogResp struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	StrainID    int       `json:"strain_id"`
}
