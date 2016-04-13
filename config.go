package main

import (
	"github.com/ddollar/forego/Godeps/_workspace/src/github.com/subosito/gotenv"
	"os"
)

type Config map[string]string

func ReadConfig(filename string) (Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return make(Config), nil
	}
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	config := make(Config)
	for key, val := range gotenv.Parse(fd) {
		config[key] = val
	}
	return config, nil
}
