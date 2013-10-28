package server

import (
	"testing"
	"os"
	"path"
	"strings"
)

var (
	directory, _ = os.Getwd()
	templatePaths = []string{path.Join(directory, "fixtures/views")}
)

func TestTemploadLoadingPaths(t *testing.T) {
	templateLoader := NewTemplateLoader(templatePaths)
	templateLoader.Refresh()

	if len(templateLoader.templates) == 0 {
		t.Errorf("Could not find templates under fixtures/views")
	}

	if _, ok := templateLoader.templates["Index.html"] ; !ok {
		t.Errorf("Could not find template Index.html")
	}
}

func TestTemplateProperties(t *testing.T) {
	templateLoader := NewTemplateLoader(templatePaths)
	templateLoader.Refresh()

	template := templateLoader.templates["Index.html"]

	if template.Name() != "Index.html" {
		t.Errorf("Could not resolve the name of the template")
	}

	content := template.Content()
	if !strings.Contains(content, "<html><body></body></html>") {
		t.Errorf("Could not read content of the template")
	}
}

func TestHtmlRender(t *testing.T) {
	templateLoader := NewTemplateLoader(templatePaths)
	templateLoader.Refresh()
	// TODO: write test for testing render implementation here.
}

