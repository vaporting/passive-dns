package passivehandler

import (
	"github.com/go-pg/pg"

	"passive-dns/types"

	"fmt"
)

// IPassiveHandler is base interface for polymorphism uses
type IPassiveHandler interface {
	Refresh(entry types.ResolvedEntry, sourceID uint) error
	mergeOverlap(row *types.ResolvedRow) ([]uint, error)
	makeRow(entry types.ResolvedEntry, sourceID uint) (types.ResolvedRow, error)
	reorganizeResolutions(row *types.ResolvedRow, overlaps []uint) error
}

type passiveHandler struct {
	ResolvedTypes []string
	whereFmtCmd   string
	db            *pg.DB

	// Interface
	IPassiveHandler
}

// Refresh refreshes the rows in table:resolved_[element]
func (handler *passiveHandler) Refresh(entry types.ResolvedEntry, sourceID uint) error {
	fmt.Println("passivehander in")
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

func (handler *passiveHandler) mergeOverlap(row *types.ResolvedRow) ([]uint, error) {
	return []uint{}, nil
}
func (handler *passiveHandler) makeRow(entry types.ResolvedEntry, sourceID uint) (types.ResolvedRow, error) {
	return types.ResolvedRow{}, nil
}

func (handler *passiveHandler) reorganizeResolutions(row *types.ResolvedRow, overlaps []uint) error {
	return nil
}
