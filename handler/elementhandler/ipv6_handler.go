package elementhandler

import (
	"fmt"

	"strings"

	"github.com/jinzhu/gorm"

	"passive-dns/types"

	"passive-dns/db"
)

// Ipv6Handler is used to handler Ipv6 entry from request
type Ipv6Handler struct {
	Type              string
	db                *gorm.DB
	prefixInsertCmd   string
	ignoreConflictCmd string
}

// Refresh updates table:ips
func (handler *Ipv6Handler) Refresh(entries []string) error {
	fmt.Println("this is ipv6Handler func:refresh")
	if len(entries) == 0 {
		return nil
	}
	cmd := handler.prefixInsertCmd
	cmd += fmt.Sprintf(" ('%s', '%s', NOW(), NOW())", entries[0], handler.Type)
	for _, entry := range entries[1:] {
		entry = strings.Replace(entry, "'", "''", -1) // avoid sql injection
		cmd += fmt.Sprintf(", ('%s', '%s', NOW(), NOW())", entry, handler.Type)
	}
	cmd += handler.ignoreConflictCmd
	// fmt.Println(cmd)
	_, err := handler.db.CommonDB().Exec(cmd)
	return err
}

// NewIpv6Handler creates Ipv6Handler with certain type
func NewIpv6Handler() *Ipv6Handler {
	handler := Ipv6Handler{}
	tempDB, _ := db.GetDB()
	handler.db = tempDB
	handler.Type = types.DNSIpv6Type
	handler.prefixInsertCmd = "INSERT INTO ips (ip, type, created_at, updated_at) VALUES"
	handler.ignoreConflictCmd = " ON CONFLICT DO NOTHING;"
	return &handler
}
