package scaffold

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jnawaz/kmp-gen/internal/config"
)

//go:embed templates
var templateFS embed.FS

type templateData struct {
	Package          string
	FeatureName      string
	FeatureNameTitle string
}

func Feature(featureName string, cfg config.Config) error {
	data := templateData{
		Package:          cfg.Package,
		FeatureName:      featureName,
		FeatureNameTitle: title(featureName),
	}

	var files []FileSpec
	switch cfg.Architecture {
	default:
		files = dddFiles(cfg.SourceRoot, featureName, cfg.Package)
	}

	for _, f := range files {
		if err := generateFile(f, data); err != nil {
			return err
		}
	}

	return nil
}

func generateFile(f FileSpec, data templateData) error {
	tmplPath := filepath.Join("templates", "ddd", f.TemplateName)
	tmplContent, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(f.TemplateName).Parse(string(tmplContent))
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(f.OutputPath), 0755); err != nil {
		return err
	}

	out, err := os.Create(f.OutputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, data)
}
