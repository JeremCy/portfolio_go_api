package models

import "gorm.io/datatypes"

type Project struct {
	ID     uint           `json:"id" gorm:"primary_key"`
	Title  string         `json:"title"`
	Desc   string         `json:"desc"`
	Descfr string         `json:"descfr"`
	Stacks datatypes.JSON `json:"stacks"`
}
