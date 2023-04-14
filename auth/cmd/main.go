package main

import (
	"log"
)

func main() {
	if err := GetRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
	// ...
}
