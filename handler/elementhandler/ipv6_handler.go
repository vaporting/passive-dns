package elementhandler

import (
	"fmt"

	"passive-dns/types"
)

// Ipv6Handler is used to handler Ipv6 entry from request
type Ipv6Handler struct {
	*ipHandler
}

// Refresh updates table:ips
func (handler *Ipv6Handler) Refresh(entries []string) error {
	fmt.Println("this is ipv6Handler func:refresh")
	return handler.ipHandler.Refresh(entries)
}

// NewIpv6Handler creates Ipv6Handler with certain type
func NewIpv6Handler() *Ipv6Handler {
	handler := Ipv6Handler{ipHandler: newIPHandler()}
	handler.Type = types.DNSIpv6Type
	return &handler
}
