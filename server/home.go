package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	html "html/template"
	"github.com/shaoshing/train"
)

type Home struct {
	template *html.Template
}

func (home *Home) Initialize() {
	files := [2]string {
				"views/layout.html",
				"views/index.html",
			}

	home.template =
		html.New("index").
			Funcs(train.HelperFuncs)

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("[Error] " + err.Error())
		}
		home.template.Parse(string(content))
	}
}

func (home *Home) Index (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	args := []string{}
	if err := home.template.Execute(response, args) ; err != nil {
		http.Error(response, err.Error(), 400)
	}
}

