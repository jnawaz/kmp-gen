package scaffold

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jnawaz/kmp-gen/internal/config"
)

func TestFeature(t *testing.T) {
	t.Run("scaffolds all ddd files", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		cfg := config.Config{
			Package:      "com.example.app",
			SourceRoot:   "src/commonMain/kotlin",
			Architecture: "ddd",
		}

		if err := Feature("auth", cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		base := filepath.Join(dir, "src/commonMain/kotlin", "com/example/app", "feature", "auth")

		expectedFiles := []string{
			filepath.Join(base, "domain", "model", "Auth.kt"),
			filepath.Join(base, "domain", "repository", "AuthRepository.kt"),
			filepath.Join(base, "domain", "usecase", "GetAuthUseCase.kt"),
			filepath.Join(base, "data", "repository", "AuthRepositoryImpl.kt"),
			filepath.Join(base, "data", "source", "AuthRemoteSource.kt"),
			filepath.Join(base, "data", "source", "AuthLocalSource.kt"),
			filepath.Join(base, "presentation", "AuthViewModel.kt"),
			filepath.Join(base, "presentation", "AuthScreen.kt"),
		}

		for _, path := range expectedFiles {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				t.Errorf("expected file not created: %s", path)
			}
		}
	})

	t.Run("does not overwrite existing feature", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}

		cfg := config.Config{
			Package:      "com.example.app",
			SourceRoot:   "src/commonMain/kotlin",
			Architecture: "ddd",
		}

		if err := Feature("auth", cfg); err != nil {
			t.Fatalf("unexpected error scaffolding auth: %v", err)
		}
		if err := Feature("payments", cfg); err != nil {
			t.Fatalf("unexpected error scaffolding payments: %v", err)
		}

		base := filepath.Join(dir, "src/commonMain/kotlin", "com/example/app", "feature")

		for _, name := range []string{"auth", "payments"} {
			path := filepath.Join(base, name, "domain", "model", title(name)+".kt")
			if _, err := os.Stat(path); os.IsNotExist(err) {
				t.Errorf("expected feature %q to exist at %s", name, path)
			}
		}
	})
}
