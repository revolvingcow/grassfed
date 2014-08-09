package models

import (
	"time"
)

type History struct {
	Id        int64
	AccountId int64
	Product   string
	Calories  int64
	Date      time.Time
}
