package models

import (
    "time"
)

type Account struct {
    Id              int64
    Profile         string
    Created         time.Time
    LastVisit       time.Time
}
