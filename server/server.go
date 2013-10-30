package server

import (
	"net/http"
	"fmt"
)

func Server() {
	fmt.Println("Loading controllers...")
	controllers()

	fmt.Println("Listen and serve...")
	http.ListenAndServe(":9002", nil)
}

func controllers() {
	Setup()
}

