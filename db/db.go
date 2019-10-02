package db

import (
	"github.com/go-pg/pg"

	"errors"
)

var passiveDB *pg.DB

// InitDB initializes the database only once
func InitDB() (*pg.DB, error) {
	var err error = nil
	if passiveDB == nil {
		passiveDB = pg.Connect(&pg.Options{
			User:     "",
			Password: "",
			Database: "passivedns_dev",
		})
		if _, err = passiveDB.Exec("SELECT 1;"); err != nil {
			passiveDB = nil
		}
	}
	return passiveDB, err
}

// GetDB get DB object, if DB is initialized
func GetDB() (*pg.DB, error) {
	err := errors.New("DB is uninitialized")
	if passiveDB != nil {
		err = nil
	}
	return passiveDB, err
}

// CloseDB closes DB
func CloseDB() {
	if passiveDB != nil {
		passiveDB.Close()
	}
}
