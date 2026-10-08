package scaffold

import (
	"path/filepath"
	"strings"
)

type FileSpec struct {
	TemplateName string
	OutputPath   string
}

func packageToPath(packageName string) string {
	return strings.ReplaceAll(packageName, ".", "/")
}

func dddFiles(sourceRoot, featureName, packageName string) []FileSpec {
	base := filepath.Join(sourceRoot, packageToPath(packageName), "feature", featureName)

	return []FileSpec{
		{
			TemplateName: "domain/model.kt.tmpl",
			OutputPath:   filepath.Join(base, "domain", "model", title(featureName)+".kt"),
		},
		{
			TemplateName: "domain/repository.kt.tmpl",
			OutputPath:   filepath.Join(base, "domain", "repository", title(featureName)+"Repository.kt"),
		},
		{
			TemplateName: "domain/usecase.kt.tmpl",
			OutputPath:   filepath.Join(base, "domain", "usecase", "Get"+title(featureName)+"UseCase.kt"),
		},
		{
			TemplateName: "data/repository_impl.kt.tmpl",
			OutputPath:   filepath.Join(base, "data", "repository", title(featureName)+"RepositoryImpl.kt"),
		},
		{
			TemplateName: "data/remote_source.kt.tmpl",
			OutputPath:   filepath.Join(base, "data", "source", title(featureName)+"RemoteSource.kt"),
		},
		{
			TemplateName: "data/local_source.kt.tmpl",
			OutputPath:   filepath.Join(base, "data", "source", title(featureName)+"LocalSource.kt"),
		},
		{
			TemplateName: "presentation/viewmodel.kt.tmpl",
			OutputPath:   filepath.Join(base, "presentation", title(featureName)+"ViewModel.kt"),
		},
		{
			TemplateName: "presentation/screen.kt.tmpl",
			OutputPath:   filepath.Join(base, "presentation", title(featureName)+"Screen.kt"),
		},
	}
}

func title(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
