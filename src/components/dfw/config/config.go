package config

import (
	"errors"
	"os"
)


func ReadCfg(cfg string) ([]string, error) {
    fd, err := os.Open(cfg)
    defer fd.Close()
    if err != nil {
		return []string{}, errors.New("cannot open config file")
    }

	return []string{}, nil
}
