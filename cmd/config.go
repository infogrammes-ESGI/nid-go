package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ListenInterface string `yaml:"listen_interface"`
}

func ParseConfigFile(path string) (Config, error) {
	var res Config

	file, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer file.Close()

	file_content, err := io.ReadAll(file)
	if err != nil {
		return res, err
	}

	err = yaml.Unmarshal(file_content, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
