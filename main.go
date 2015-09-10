package main

import (
	"flag"
	"log"
)

func main() {
	log.Println("Starting server.")
	// define required flags
	var path string
	flag.StringVar(&path, "path", "temp", "File directory path in which to run the server.")
	// parse flags
	flag.Parse()
	// TODO check & ask whether to create if none currently exists (also add flag for this?)
	// TODO if path wasn't given, ask for it (see shared code from tinzenite/tin)
	log.Println("Path:", path)
}
