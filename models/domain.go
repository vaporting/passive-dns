package models

// Domain is the table:domains template
type Domain struct {
	BaseModel
	Name string `gorm:"type:text;unique;not null"`
}