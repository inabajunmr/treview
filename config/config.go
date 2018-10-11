package config

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// Config for treview
type config struct {
	Lang []string `yaml:"lang"`
}

// Read config
func read(path string) config {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		print("Config is something wrong.")
		os.Exit(1)
	}

	var d config
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		print("Config is something wrong.")
		os.Exit(1)
	}

	return d
}

// Write config
func write(path string, conf config) {
	buf, err := yaml.Marshal(conf)
	if err != nil {
		print("Failed to write config.")
		os.Exit(1)
	}

	if err = ioutil.WriteFile(path, buf, 0666); err != nil {
		print("Failed to write config.")
		os.Exit(1)
	}
}

// SetLangs to config
func SetLangs(path string, langs []string) {
	conf := config{Lang: langs}
	write(path, conf)
}

// GetLangs from config
func GetLangs(path string) []string {
	return read(path).Lang
}
