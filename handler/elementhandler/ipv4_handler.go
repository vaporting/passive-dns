package elementhandler

import (
	// "fmt"

	"passive-dns/types"
)

// Ipv4Handler is used to handler Ipv4 entry from request
type Ipv4Handler struct {
	*ipHandler
}

// Refresh updates table:ips
func (handler *Ipv4Handler) Refresh(entries []string) error {
	// fmt.Println("this is ipv4Handler func:refresh")
	return handler.ipHandler.Refresh(entries)
}

// NewIpv4Handler creates Ipv4Handler with certain type
func NewIpv4Handler() *Ipv4Handler {
	handler := Ipv4Handler{ipHandler: newIPHandler()}
	handler.Type = types.DNSIpv4Type
	return &handler
}
