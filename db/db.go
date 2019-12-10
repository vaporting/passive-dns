package db

import (
	"github.com/go-pg/pg"

	"errors"

	"passive-dns/types"
)

var passiveDB *pg.DB

// InitDB initializes the database only once
func InitDB(config *types.Config) (*pg.DB, error) {
	var err error = nil
	if passiveDB == nil {
		passiveDB = pg.Connect(&pg.Options{
			Addr:     config.DB.Host + ":" + config.DB.Port,
			User:     config.DB.User,
			Password: config.DB.PWD,
			Database: config.DB.Name,
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
