package models

import "time"

type Goal struct {
	Id        int64
	AccountId int64
	Calories  int64
	Date      time.Time
}
