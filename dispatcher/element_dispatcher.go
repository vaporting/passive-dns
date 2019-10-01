package dispatcher

import (
	"passive-dns/handler/elementhandler"

	"passive-dns/types"
)

// ElementDispatcher dispatchs the element entry to the right handler
type ElementDispatcher struct {
	handler map[string]elementhandler.ElementHandler
}

// Init intializes the ElementDispatcher
func (dispatcher *ElementDispatcher) Init() {
	dispatcher.handler = make(map[string]elementhandler.ElementHandler)
	ipv4Handler := elementhandler.NewIpv4Handler()
	dispatcher.handler[ipv4Handler.Type] = ipv4Handler
	ipv6Handler := elementhandler.NewIpv6Handler()
	dispatcher.handler[ipv6Handler.Type] = ipv6Handler
	domainHandler := elementhandler.NewDomainHandler()
	dispatcher.handler[domainHandler.Type] = domainHandler
}

// Refresh update the db with input elements
func (dispatcher *ElementDispatcher) Refresh(source string, resolvedEntries []types.ResolvedEntry) error {
	groups := dispatcher.groupElements(resolvedEntries)
	for dnsType, group := range groups {
		err := dispatcher.handler[dnsType].Refresh(group)
		if err != nil {
			return err
		}
	}
	return nil
}

// groupElements sort the entries to its' group
func (dispatcher *ElementDispatcher) groupElements(resolvedEntries []types.ResolvedEntry) map[string][]string {
	groups := make(map[string][]string)
	for key := range dispatcher.handler {
		groups[key] = []string{}
	}

	for _, resolvedEntry := range resolvedEntries {
		_, ok := groups[resolvedEntry.Type]
		if ok {
			groups[types.DNSDomainType] = append(groups[types.DNSDomainType], resolvedEntry.Name)
			groups[resolvedEntry.Type] = append(groups[resolvedEntry.Type], resolvedEntry.Value)
		}
	}
	return groups
}
