package main

import (
	"github.com/BurntSushi/toml"
)

type Settings struct {
	Srv string
	Log string
	Db  string
}

func LoadConfig(confPath string) (Settings, error) {
	c := Settings{}
	_, err := toml.DecodeFile(confPath, &c)

	return c, err
}
