package main

import (
	"passive-dns/routes"

	"fmt"

	"passive-dns/cache"

	"passive-dns/util"
)

func main() {
	// parse configuration
	config, err := util.ReadConfig()
	if err != nil {
		return
	}
	// initialize cacher
	_, err = cache.CreateCacher(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := routes.InitRoutes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
