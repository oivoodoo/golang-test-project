package main

import (
	"fmt"
	MusicServer "github.com/oivoodoo/music/server"
)

func main() {
	defer func() {
		fmt.Println("Panic!")

		for _, message := range MusicServer.ERROR.Messages {
			fmt.Println("Error: " + message)
		}
	}()

	fmt.Println("Starting server...")
	MusicServer.Server()
	fmt.Println("Shutdown...")
}

