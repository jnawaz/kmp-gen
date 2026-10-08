package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const ConfigFileName = "kmp-gen.yaml"

func Write(config Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFileName, data, 0o644)
}
