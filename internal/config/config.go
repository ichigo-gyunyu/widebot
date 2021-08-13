package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Token  string `yaml:"token"`
	Prefix string `yaml:"prefix"`
}

func ParseYAMLConfig(filename string) (c *Config, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	c = new(Config) // c = &Config{}
	err = yaml.NewDecoder(f).Decode(c)

	return
}
