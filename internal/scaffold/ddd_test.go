package scaffold

import (
	"path/filepath"
	"testing"
)

func TestTitle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"auth", "Auth"},
		{"payments", "Payments"},
		{"myFeature", "MyFeature"},
		{"", ""},
	}

	for _, tt := range tests {
		result := title(tt.input)
		if result != tt.expected {
			t.Errorf("title(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestPackageToPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"com.example.app", "com/example/app"},
		{"com.sightline.sightlinekmp", "com/sightline/sightlinekmp"},
	}

	for _, tt := range tests {
		result := packageToPath(tt.input)
		if result != tt.expected {
			t.Errorf("packageToPath(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestDddFiles(t *testing.T) {
	files := dddFiles("composeApp/src/commonMain/kotlin", "auth", "com.example.app")

	if len(files) != 8 {
		t.Fatalf("expected 8 files, got %d", len(files))
	}

	base := filepath.Join("composeApp/src/commonMain/kotlin", "com/example/app", "feature", "auth")

	expected := []string{
		filepath.Join(base, "domain", "model", "Auth.kt"),
		filepath.Join(base, "domain", "repository", "AuthRepository.kt"),
		filepath.Join(base, "domain", "usecase", "GetAuthUseCase.kt"),
		filepath.Join(base, "data", "repository", "AuthRepositoryImpl.kt"),
		filepath.Join(base, "data", "source", "AuthRemoteSource.kt"),
		filepath.Join(base, "data", "source", "AuthLocalSource.kt"),
		filepath.Join(base, "presentation", "AuthViewModel.kt"),
		filepath.Join(base, "presentation", "AuthScreen.kt"),
	}

	for i, f := range files {
		if f.OutputPath != expected[i] {
			t.Errorf("file[%d] path = %q, want %q", i, f.OutputPath, expected[i])
		}
	}
}
