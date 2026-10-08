package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectSourceRoot(t *testing.T) {
	t.Run("find single commonMain/kotlin", func(t *testing.T) {
		dir := t.TempDir()
		os.MkdirAll(filepath.Join(dir, "composeApp/src/commonMain/kotlin"), 0o755)

		err := os.Chdir(dir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		result, err := DetectSourceRoot()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == "" {
			t.Error("expected a path, got empty string")
		}
	})

	t.Run("returns empty when no match", func(t *testing.T) {
		dir := t.TempDir()
		err := os.Chdir(dir)
		if err != nil {
			t.Fatalf("failed to change directory %v", err)
		}

		result, err := DetectSourceRoot()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("returns empty when multiple matches", func(t *testing.T) {
		dir := t.TempDir()
		os.MkdirAll(filepath.Join(dir, "composeApp/src/commonMain/kotlin"), 0o755)
		os.MkdirAll(filepath.Join(dir, "shared/src/commonMain/kotlin"), 0o755)

		if err := os.Chdir(dir); err != nil {
			t.Fatalf("failed to change directory %v", err)
		}

		result, err := DetectSourceRoot()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != "" {
			t.Errorf("unexpected empty string for ambiguous result, got %q", result)
		}
	})
}
