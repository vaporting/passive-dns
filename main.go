package main

import (
	"passive-dns/routes"

	"fmt"

	"os"

	"passive-dns/db"

	"passive-dns/types"

	"gopkg.in/yaml.v2"
)

func main() {
	// parse configuration
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	config := types.Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	// initialize database
	db.InitDB(&config)
	_, err = db.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := routes.InitRoutes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
