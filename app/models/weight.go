package models

import "time"

type Weight struct {
    Id              int64
    AccountId       int64
    Weight          float64
    Date            time.Time
}
