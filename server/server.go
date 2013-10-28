package server

import (
	"net/http"
	"os"
	"path"
	"fmt"
)

var (
	MainTemplateLoader *TemplateLoader
	TemplatePaths []string
)

func Server() {
	fmt.Println("Loading paths...")
	paths()
	fmt.Println("Loading templates...")
	templates()
	fmt.Println("Loading controllers...")
	controllers()

	fmt.Println("Loading views from the default directories...")
	MainTemplateLoader.Refresh()

	fmt.Println("Listen and serve...")
	// Here we are setuping all our controllers and assets pipeline.
	http.ListenAndServe(":9001", nil)
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
	Setup()
}

