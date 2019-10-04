package hunter

import (
	"github.com/go-pg/pg"
)

// IHunter is base interface for polymorphism uses
type IHunter interface {
	Hunt(sources []string) ([]byte, error)
}

type hunter struct {
	SourceTypes []string
	joinFmtCmd  string
	db          *pg.DB

	// Interface
	IHunter
}
