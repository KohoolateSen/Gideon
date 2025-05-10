package utils

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func JSONParse(path string, val any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, val); err != nil {
		return err
	}

	return nil
}

func YAMLParse(path string, val any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, val); err != nil {
		return err
	}

	return nil
}
