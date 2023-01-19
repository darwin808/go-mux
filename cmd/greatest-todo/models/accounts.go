package models

type Accounts struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Todo       string `json:"todo"`
	Created_On string `json:"created_on"`
}
