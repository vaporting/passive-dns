package passivehandler

import (
	"github.com/go-pg/pg"

	"time"

	"passive-dns/types"

	"passive-dns/models"

	"passive-dns/db"
)

// PassiveIPHandler handles the resolved ip entry from request
type PassiveIPHandler struct {
	*passiveHandler
}

// Refresh refreshes the rows in table:resolved_[element]
func (handler *PassiveIPHandler) Refresh(entry types.ResolvedEntry, sourceID uint) error {
	overlaps := []uint{}
	newRow, err := handler.makeRow(entry, sourceID)
	if err == nil {
		overlaps, err = handler.mergeOverlap(&newRow)
	}

	if err == nil {
		err = handler.reorganizeResolutions(&newRow, overlaps)
	}
	return err
}

// mergeOverlap merge overlap rows and new row to new row. Return overlap row's id
func (handler *PassiveIPHandler) mergeOverlap(row *types.ResolvedRow) ([]uint, error) {
	overlaps := []models.ResolvedIP{}
	err := handler.db.Model(&overlaps).
		Where(handler.whereFmtCmd, row.OriginalID, row.PassiveID, row.SourceID, row.FirstSeen, row.LastSeen).
		Order("first_seen ASC").
		Select()
	ids := make([]uint, len(overlaps))
	if err == nil && len(overlaps) > 0 {
		// because overlap_entries are sorted through first_seen from oldest to latest,
		// we just need to check the first first_seen and the last last_seen
		if row.FirstSeen.After(overlaps[0].FirstSeen) {
			row.FirstSeen = overlaps[0].FirstSeen
		}
		if row.LastSeen.Before(overlaps[len(overlaps)-1].LastSeen) {
			row.LastSeen = overlaps[len(overlaps)-1].LastSeen
		}

		// extract overlap id
		for index, overlap := range overlaps {
			ids[index] = overlap.ID
		}
	}
	return ids, err
}

// makeRow make new row with resolved entry from request
func (handler *PassiveIPHandler) makeRow(entry types.ResolvedEntry, sourceID uint) (types.ResolvedRow, error) {
	domain := models.Domain{}
	ip := models.IP{}
	firstSeen := time.Now()
	lastSeen := firstSeen
	err := handler.db.Model(&domain).Where("name = ?", entry.Name).Select()
	if err == nil {
		err = handler.db.Model(&ip).Where("ip = ?", entry.Value).Select()
	}
	if err == nil {
		firstSeen, err = time.Parse("2006-01-02", entry.FirstSeen)
	}
	if err == nil {
		lastSeen, err = time.Parse("2006-01-02", entry.LastSeen)
	}
	return types.ResolvedRow{
		OriginalID: domain.ID,
		PassiveID:  ip.ID,
		SourceID:   sourceID,
		FirstSeen:  firstSeen,
		LastSeen:   lastSeen}, err
}

// reorganizeResolutions deletes overlap rows and inserts new Row. Actions are wrapped by transcation
func (handler *PassiveIPHandler) reorganizeResolutions(row *types.ResolvedRow, overlaps []uint) error {
	err := handler.db.RunInTransaction(
		func(tx *pg.Tx) error {
			var err error = nil
			if len(overlaps) > 0 {
				_, err = tx.Model((*models.ResolvedIP)(nil)).Where("id IN (?)", pg.In(overlaps)).Delete()
				if err != nil {
					return err
				}
			}
			_, err = tx.Model(&models.ResolvedIP{
				DomainID:     row.OriginalID,
				ResolvedIpID: row.PassiveID,
				SourceID:     row.SourceID,
				FirstSeen:    row.FirstSeen,
				LastSeen:     row.LastSeen,
				BaseModel:    models.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}}).Insert()
			return err
		})
	return err
}

// NewPassiveIPHandler creates PassiveIPHandler
func NewPassiveIPHandler() *PassiveIPHandler {
	handler := PassiveIPHandler{passiveHandler: &passiveHandler{}}
	tempDB, _ := db.GetDB()
	handler.db = tempDB
	handler.ResolvedTypes = []string{types.DNSIpv4Type, types.DNSIpv6Type}
	handler.whereFmtCmd = "domain_id = ? AND resolved_ip_id = ? AND source_id = ? AND ? <= last_seen AND ? >= first_seen"
	return &handler
}
