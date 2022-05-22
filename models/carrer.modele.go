package models

import "time"

type Carrer struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	School   string    `json:"title"`
	Desc     string    `json:"desc"`
	Descfr   string    `json:"descfr"`
	Start_at time.Time `json:"start_at"`
	End_at   time.Time `json:"end_at"`
}
