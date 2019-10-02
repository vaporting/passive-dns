package elementhandler

import (
	// "fmt"

	"time"

	"passive-dns/types"

	"passive-dns/db"

	"passive-dns/models"
)

// ipHandler is used to handler Ipv4 entry from request
type ipHandler struct {
	*baseElementHandler
}

// Refresh updates table:ips
func (handler *ipHandler) Refresh(entries []string) error {
	// fmt.Println("this is ipHandler func:refresh")
	if len(entries) == 0 {
		return nil
	}
	ips := []models.IP{}
	for _, entry := range entries {
		ips = append(
			ips,
			models.IP{
				BaseModel: models.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()},
				IP:        entry,
				Type:      handler.Type})
	}
	_, err := handler.db.Model(&ips).OnConflict("DO NOTHING").Insert()
	return err
}

// newIPHandler creates ipHandler with certain type
func newIPHandler() *ipHandler {
	handler := ipHandler{baseElementHandler: &baseElementHandler{}}
	tempDB, _ := db.GetDB()
	handler.db = tempDB
	handler.Type = types.DNSIpv4Type
	return &handler
}
