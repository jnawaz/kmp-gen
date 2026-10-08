// Package config
package config

type Config struct {
	Package      string `yaml:"package"`
	SourceRoot   string `yaml:"sourceRoot"`
	Architecture string `yaml:"architecture"`
}
