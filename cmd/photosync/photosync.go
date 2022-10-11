package main

import (
	"log"
	"os"

	"github.com/photoprism/photoprism/internal/photoprism"
)

func main() {
	f, err := photoprism.NewMediaFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	log.Printf("name: %s", f.CanonicalName())
}
