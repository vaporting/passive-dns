package elementhandler

import (
	"fmt"

	"strings"

	"github.com/jinzhu/gorm"

	"passive-dns/types"

	"passive-dns/db"
)

// DomainHandler is used to handler domain entry from request
type DomainHandler struct {
	Type              string
	db                *gorm.DB
	prefixInsertCmd   string
	ignoreConflictCmd string
}

// Refresh updates table:domains
func (handler *DomainHandler) Refresh(entries []string) error {
	fmt.Println("this is domainHandler func:refresh")
	if len(entries) == 0 {
		return nil
	}
	cmd := handler.prefixInsertCmd
	cmd += fmt.Sprintf(" ('%s', NOW(), NOW())", entries[0])
	for _, entry := range entries[1:] {
		entry = strings.Replace(entry, "'", "''", -1) // avoid sql injection
		cmd += fmt.Sprintf(", ('%s', NOW(), NOW())", entry)
	}
	cmd += handler.ignoreConflictCmd
	// fmt.Println(cmd)
	_, err := handler.db.CommonDB().Exec(cmd)
	return err
}

// NewDomainHandler creates DomainHandler with certain type
func NewDomainHandler() *DomainHandler {
	handler := DomainHandler{}
	tempDB, _ := db.GetDB()
	handler.db = tempDB
	handler.Type = types.DNSDomainType
	handler.prefixInsertCmd = "INSERT INTO domains (name, created_at, updated_at) VALUES"
	handler.ignoreConflictCmd = " ON CONFLICT DO NOTHING;"
	return &handler
}
