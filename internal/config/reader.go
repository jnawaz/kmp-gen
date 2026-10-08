package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Read() (Config, error) {
	data, err := os.ReadFile(ConfigFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, fmt.Errorf("kmp-gen.yaml not found — run `kmp-gen init` first")
		}
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
