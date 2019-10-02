package elementhandler

import "github.com/go-pg/pg"

// IElementHandler is base interface for polymorphism uses
type IElementHandler interface {
	Refresh(entries []string) error
}

// baseElementHandler is abstract struct.
type baseElementHandler struct {
	Type string
	db   *pg.DB

	// Interface
	IElementHandler
}

func (handler *baseElementHandler) Refresh(entries []string) error {
	// embedding struct needs to implement
	return nil
}
