package models

import "time"

// BaseModel is the base model for all db table
type BaseModel struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
