package util

import (
	"os"

	"fmt"

	"passive-dns/types"

	"gopkg.in/yaml.v2"

	"github.com/kelseyhightower/envconfig"
)

// ReadConfig reads config from env or config.yml
func ReadConfig() (*types.Config, error) {
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
