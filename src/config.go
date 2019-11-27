package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// ---------------------------------------------------------------------------------------------------------------------

type config struct {
	PGMockBindAddress  string `yaml:"pgmock-bind-addr"`
	DLoaderBindAddress string `yaml:"dloader-bind-addr"`
}

// ---------------------------------------------------------------------------------------------------------------------

func loadConfig(configPath string) (*config, error) {

	// create the blank config
	cfg := &config{}

	// if the path is empty return fail
	if len(configPath) == 0 {
		return nil, fmt.Errorf("unable to read config file, path is empty")
	}

	// attempt to read in the config file bytes
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// attempt to unmarshal the JSON
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}

	// all good, return the config
	return cfg, nil
}
