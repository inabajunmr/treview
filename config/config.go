package config

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// Config for treview
type Config struct {
	Lang []string `yaml:"lang"`
}

// Read config
func Read(path string) Config {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		print("Config is something wrong.")
		os.Exit(1)
	}

	var d Config
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		print("Config is something wrong.")
		os.Exit(1)
	}

	return d
}
