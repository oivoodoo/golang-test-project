package server

import "net/http"

func Setup() {
	home := new(Home)
	http.HandleFunc("/", home.Index)
}

