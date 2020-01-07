package main

import (
	"passive-dns/routes"

	"fmt"

	"os"

	"passive-dns/db"

	"passive-dns/types"

	"gopkg.in/yaml.v2"

	"github.com/kelseyhightower/envconfig"
)

func readConfig() (*types.Config, error) {
	config := types.Config{}
	// parse from Env
	err := envconfig.Process("DB", &config)
	if err != nil {
		fmt.Println(err)
		// parse from file
		var f *os.File
		f, err = os.Open("config.yml")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		config = types.Config{}
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&config)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &config, err
}

func main() {
	// parse configuration
	config, err := readConfig()
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
