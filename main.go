package main

import (
	"fmt"
	MusicServer "github.com/oivoodoo/music/server"
)

func main() {
	defer func() {
		fmt.Println("Panic!")
	}()

	fmt.Println("Starting server...")
	MusicServer.Server()
	fmt.Println("Shutdown...")
}
