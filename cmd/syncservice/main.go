package main

import (
	"fmt"

	"passive-dns/cache"

	"passive-dns/db"

	"passive-dns/util"

	"passive-dns/syncservice/synchronizer"
)

func main() {
	config, err := util.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.InitDB(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = cache.CreateCacher(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	syncers := synchronizer.CreateSyncers(config)
	// execute synchronous process
	for i := 0; i < len(syncers); i++ {
		syncers[i].Sync()
	}
	db.CloseDB()
	cache.CloseCacher()
}
