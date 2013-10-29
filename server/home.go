package server

import (
	"net/http"
)

type Home struct {}

func (home *Home) Index (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	args := []string{}
	MainTemplateLoader.templates["Index.html"].Render(response, args)
}

