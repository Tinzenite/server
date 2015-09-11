package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/tinzenite/encrypted"
	"github.com/tinzenite/shared"
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

	if exists, _ := shared.DirectoryExists(path); !exists {
		shared.MakeDirectory(path)
	}

	enc, err := encrypted.Create(path, "d_server")
	if err != nil {
		log.Println("Server: failed to create:", err)
		return
	}
	// print important info
	address, _ := enc.Address()
	fmt.Printf("Running peer <%s>.\nID: %s\n", enc.Name(), address)
	// prepare quitting via ctrl-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// loop until close
	for {
		select {
		case <-c:
			log.Println("Server: quitting.")
			return
		} // select
	} // for
}
