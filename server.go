package music

import (
	"net/http"
	"os"
	"path"
	Controllers "github.com/oivoodoo/music/controllers"
)

var (
	MainTemplateLoader *TemplateLoader
	TemplatePaths []string
)

func Server() {
	paths()
	templates()
	controllers()

	MainTemplateLoader.Refresh()

	// Here we are setuping all our controllers and assets pipeline.
	http.ListenAndServe(":9000", nil)
}

func paths() {
	directory, _ := os.Getwd()
	TemplatePaths = append(TemplatePaths, path.Join(directory, "views"))
	TemplatePaths = append(TemplatePaths, path.Join(directory, "api/views"))
}

func templates() {
	MainTemplateLoader = NewTemplateLoader(TemplatePaths)
}

func controllers() {
	Controllers.Setup()
}

