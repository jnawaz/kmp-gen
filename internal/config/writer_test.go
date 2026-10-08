package config

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestWrite(t *testing.T) {
	t.Run("creates kmp-gen.yaml with correct content", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("failed to change directory %v", err)
		}

		config := Config{
			Package:      "com.example.app",
			SourceRoot:   "composeApp/src/commonMain/kotlin",
			Architecture: "ddd",
		}

		if err := Write(config); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		data, err := os.ReadFile(filepath.Join(dir, ConfigFileName))
		if err != nil {
			t.Fatalf("config file was not created %v", err)
		}

		var result Config
		if err := yaml.Unmarshal(data, &result); err != nil {
			t.Fatalf("failed to parse written yaml: %v", err)
		}

		if result.Package != config.Package {
			t.Errorf("expected package %q, got %q", config.Package, result.Package)
		}
		if result.SourceRoot != config.SourceRoot {
			t.Errorf("expected SourceRoot %q, got %q", config.SourceRoot, result.SourceRoot)
		}
		if result.Architecture != config.Architecture {
			t.Errorf("expected Architecture %q, got %q", config.Architecture, result.Architecture)
		}
	})
}
