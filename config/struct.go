package config

import "time"

type Config struct {
	Env    string
	Name   string       `yaml:"name"`
	DB     dbStruct     `yaml:"db"`
	Time   timeStruct   `yaml:"time"`
	Logger loggerStruct `yaml:"logger"`
}

type dbStruct struct {
	Type    string `yaml:"type"`
	Params  string `yaml:"params"`
	MaxIdle int    `yaml:"max_idle"`
	MaxOpen int    `yaml:"max_open"`
}

type timeStruct struct {
	Zone string `yaml:"zone"`
}

type loggerStruct struct {
	Enabled  bool   `yaml:"enabled"`
	Filename string `yaml:"filename,omitempty"`
}

func (c *Config) GetLocation() *time.Location {
	loc, err := time.LoadLocation(c.Time.Zone)
	if err != nil {
		panic(err)
	}
	return loc
}
