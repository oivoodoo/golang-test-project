package server

import (
	"github.com/shaoshing/train"
	"net/http"
)

func Setup() {
	train.ConfigureHttpHandler(nil)

	home := new(Home)
	home.Initialize()

	http.HandleFunc("/", home.Index)
}

