package config

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("reads a valid kmp-gen.yaml", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		cfg := Config{
			Package:      "com.example.app",
			SourceRoot:   "composeApp/src/commonMain/kotlin",
			Architecture: "ddd",
		}
		if err := Write(cfg); err != nil {
			t.Fatalf("failed to write config: %v", err)
		}

		result, err := Read()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.Package != cfg.Package {
			t.Errorf("expected package %q, got %q", cfg.Package, result.Package)
		}
		if result.SourceRoot != cfg.SourceRoot {
			t.Errorf("expected sourceRoot %q, got %q", cfg.SourceRoot, result.SourceRoot)
		}
		if result.Architecture != cfg.Architecture {
			t.Errorf("expected architecture %q, got %q", cfg.Architecture, result.Architecture)
		}
	})

	t.Run("returns helpful error when file missing", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		_, err := Read()
		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "kmp-gen.yaml not found — run `kmp-gen init` first"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})
}
