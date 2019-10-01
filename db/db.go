package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"errors"
)

var passiveDB *gorm.DB

// InitDB initializes the database only once
func InitDB() (*gorm.DB, error) {
	var err error = nil
	if passiveDB == nil {
		passiveDB, err = gorm.Open("postgres", "user= dbname=passivedns_dev password=")
		passiveDB.LogMode(true)
	}
	return passiveDB, err
}

// GetDB get DB object, if DB is initialized
func GetDB() (*gorm.DB, error) {
	err := errors.New("DB is uninitialized")
	if passiveDB != nil {
		err = nil
	}
	return passiveDB, err
}
