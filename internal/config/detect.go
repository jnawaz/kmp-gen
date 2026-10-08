package config

import (
	"os"
	"path/filepath"
)

func DetectSourceRoot() (string, error) {
	var found []string

	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() && d.Name() == "kotlin" {
			if filepath.Base(filepath.Dir(path)) == "commonMain" {
				found = append(found, path)
			}
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	if len(found) == 1 {
		return found[0], nil
	}

	return "", nil
}
