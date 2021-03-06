package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/tinzenite/encrypted"
	"github.com/tinzenite/shared"
)

const tag = "Main:"

func main() {
	log.Println(tag, "starting server.")
	// define required flags
	var path string
	var commandString string
	var useHadoop bool
	var address string
	var user string
	flag.StringVar(&path, "path", "", "File directory path in which to run the server.")
	flag.StringVar(&commandString, "cmd", "load", "Command for the path: create or load. Default is load.")
	flag.BoolVar(&useHadoop, "hadoop", false, "Flag to enable Hadoop storage.")
	flag.StringVar(&address, "address", "127.0.0.1", "Address of HDFS to connect to.")
	flag.StringVar(&user, "user", "root", "User of HDFS to connect as.")
	// parse flags
	flag.Parse()

	// prepare command
	command := shared.CmdParse(commandString)
	// if command isn't load or create, quit
	if command != shared.CmdLoad && command != shared.CmdCreate {
		log.Println(tag, "invalid command given!")
		return
	}
	// prepare path
	if path == "" {
		path = shared.GetString("Enter path for Server directory:")
	}
	// make sure the path is clean and absolute
	path, _ = filepath.Abs(filepath.Clean(path))
	// path may not be empty OR only contain '.' (which means error in filepath)
	if path == "" || path == "." {
		log.Println(tag, "no path given!")
		return
	}
	// create it if required and desired
	if exists, _ := shared.DirectoryExists(path); !exists {
		// if command is load we're not going to offer creating the path
		if command == shared.CmdLoad {
			log.Println(tag, "Can not run Server without valid path.")
			return
		}
		useQuestion := shared.CreateYesNo("Path <" + path + "> doesn't exist. Create it?")
		if useQuestion.Ask() < 0 {
			log.Println(tag, "Can not run Server without valid path.")
			return
		}
		// if yes, create it
		shared.MakeDirectory(path)
	}

	// prepare storage
	var store encrypted.Storage
	if useHadoop {
		var err error
		store, err = createHDFSStorage(address, user)
		if err != nil {
			log.Println("Failed to connect to HDFS:", err)
			return
		}
	} else {
		// disk storage writes data to disk
		store = createDiskStorage(path)
	}

	var enc *encrypted.Encrypted
	var err error
	switch command {
	case shared.CmdLoad:
		enc, err = encrypted.Load(path, store)
		if err != nil {
			log.Println(tag, "failed to load encrypted:", err)
			return
		}
	case shared.CmdCreate:
		peerName := shared.GetString("Please enter a peer name for this instance:")
		enc, err = encrypted.Create(path, peerName, store)
		if err != nil {
			log.Println(tag, "failed to create encrypted:", err)
			return
		}
		// store first version for future loads
		err = enc.Store()
		if err != nil {
			log.Println(tag, "failed to store initially:", err)
		}
	default:
		log.Println(tag, "No valid command was chosen, so we'll do nothing.")
		return
	}

	// run encrypted
	// print important info
	toxAddress, _ := enc.Address()
	fmt.Printf("%s Running server <%s>.\nID: %s\n", tag, enc.Name(), toxAddress)
	// prepare quitting via ctrl-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// loop until close
	for {
		select {
		case <-c:
			// store before closing
			_ = enc.Store()
			enc.Close()
			log.Println(tag, "quitting.")
			return
		} // select
	} // for
}
