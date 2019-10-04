package models

// IP is the table:ips template
type IP struct {
	BaseModel
	IP   string `gorm:"type:byte;unique;not null"`
	Type string `gorm:"type:varchar(4);not null"`
}
