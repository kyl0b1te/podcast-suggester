package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: expected command")
		os.Exit(1)
	}

	// run command
	switch os.Args[1] {
	case "cache":
		cache()
	case "audio":
		audio()
	default:
		fmt.Println("error: unknown command")
		os.Exit(1)
	}
}
