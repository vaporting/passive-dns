package elementhandler

// ElementHandler is base interface for polymorphism uses
type ElementHandler interface {
	Refresh(entries []string) error
}
