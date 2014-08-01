package models

import (
    "time"
)

type Account struct {
    Id              int64
    Profile         string
    Goal            int64
    Created         time.Time
    LastVisit       time.Time
}
