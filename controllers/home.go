package controllers

import (
	"net/http"
	"io"
	"html/template"
)

type TemplateLoader struct {
	paths []string
}

var (
	indexTemplate = template.New("Index")
)

type Home struct {}

func (home *Home) Index (response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "test")
}

