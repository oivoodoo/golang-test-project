package server

import (
	"github.com/shaoshing/train"
	"net/http"
)

func Setup() {
	train.ConfigureHttpHandler(nil)
	home := new(Home)
	http.HandleFunc("/", home.Index)
}
