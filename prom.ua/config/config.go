package config

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func New() ([]string, error) {
	cfg, err := ioutil.ReadFile("./config.txt")
	if err != nil {
		return nil, fmt.Errorf("read file failed: %v", err)
	}

	s := strings.Split(string(cfg), "\n")

	return s, nil
}
