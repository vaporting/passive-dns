package elementhandler

import (
	"fmt"

	"time"

	"passive-dns/types"

	"passive-dns/db"

	"passive-dns/models"
)

// DomainHandler is used to handler domain entry from request
type DomainHandler struct {
	*baseElementHandler
}

// Refresh updates table:domains
func (handler *DomainHandler) Refresh(entries []string) error {
	fmt.Println("this is domainHandler func:refresh")
	if len(entries) == 0 {
		return nil
	}
	domains := []models.Domain{}
	for _, entry := range entries {
		domains = append(
			domains,
			models.Domain{
				BaseModel: models.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Name:      entry,
			})
	}
	_, err := handler.db.Model(&domains).OnConflict("DO NOTHING").Insert()
	return err
}

// NewDomainHandler creates DomainHandler with certain type
func NewDomainHandler() *DomainHandler {
	handler := DomainHandler{baseElementHandler: &baseElementHandler{}}
	tempDB, _ := db.GetDB()
	handler.db = tempDB
	handler.Type = types.DNSDomainType
	return &handler
}
