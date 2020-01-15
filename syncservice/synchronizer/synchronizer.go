package synchronizer

// ISynchronizer is an interface of synchronizer
type ISynchronizer interface {
	Sync()
	extract()
	insert()
}
