package main

import (
	"passive-dns/routes"

	"fmt"

	"passive-dns/db"

	"passive-dns/util"
)

func main() {
	// parse configuration
	config, err := util.ReadConfig()
	if err != nil {
		return
	}
	// initialize database
	db.InitDB(config)
	_, err = db.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := routes.InitRoutes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
