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

// Write config
func Write(path string, conf Config) {
	buf, err := yaml.Marshal(conf)
	if err != nil {
		print("Failed to write config.")
		os.Exit(1)
	}

	if err = ioutil.WriteFile(path, buf, 6440); err != nil {
		print("Failed to write config.")
		os.Exit(1)
	}
}
