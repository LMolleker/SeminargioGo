package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// DbConfig ...
type DbConfig struct {
	Type   string `yaml:"type"`
	Driver string `yaml:"driver"`
	Conn   string `yaml:"conn"`
}

//Config ...
type Config struct {
	DbCfg   DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

//LoadConfig ...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
