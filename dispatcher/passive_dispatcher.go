package dispatcher

import (
	"passive-dns/handler/passivehandler"

	"passive-dns/types"

	"fmt"
)

// PassiveDispatcher dispatchs the element entry to the right handler
type PassiveDispatcher struct {
	handler map[string]passivehandler.IPassiveHandler
}

// Init intializes the PassiveDispatcher
func (dispatcher *PassiveDispatcher) Init() {
	dispatcher.handler = make(map[string]passivehandler.IPassiveHandler)
	passiveIPHandler := passivehandler.NewPassiveIPHandler()
	for _, rType := range passiveIPHandler.ResolvedTypes {
		dispatcher.handler[rType] = passiveIPHandler
	}
	passiveDomainHandler := passivehandler.NewPassiveDomainHandler()
	for _, rType := range passiveDomainHandler.ResolvedTypes {
		dispatcher.handler[rType] = passiveDomainHandler
	}
}

// Refresh update the db with input elements
func (dispatcher *PassiveDispatcher) Refresh(sourceID uint, resolvedEntries []types.ResolvedEntry) error {
	fmt.Println("passive dispatcher refresh")
	for _, entry := range resolvedEntries {
		err := dispatcher.handler[entry.Type].Refresh(entry, sourceID)
		fmt.Println(err)
	}
	return nil
}
