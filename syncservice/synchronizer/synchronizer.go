package synchronizer

import "passive-dns/types"

// ISynchronizer is an interface of synchronizer
type ISynchronizer interface {
	Sync()
	extract()
	insert()
}

// CreateSyncers creates synchronizers
func CreateSyncers(config *types.Config) []ISynchronizer {
	syncers := []ISynchronizer{}
	rIPSyncer, _ := NewResolvedIPSynchronizer(config)
	syncers = append(syncers, &rIPSyncer)
	return syncers
}
