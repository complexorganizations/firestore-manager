package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"
)

var (
	backup  bool
	restore bool
	path    string
)

func init() {
	if len(os.Args) > 1 {
		tempBackup := flag.Bool("backup", false, "Backup all the content of the firestore.")
		tempRestore := flag.Bool("restore", false, "Restore all the content of the firestore.")
		tempPath := flag.String("path", "example.example", "Path to the backup file.")
		flag.Parse()
		backup = *tempBackup
		restore = *tempRestore
		path = *tempPath
	} else {
		log.Fatal("No arguments provided.")
	}
	// checks
	if !backup && !restore {
		log.Fatal("No action specified.")
	}
	if path == "example.example" && filepath.Ext(path) != ".json" {
		log.Fatal("No path specified.")
	}
	if !validateJson() {
		log.Fatal("Invalid JSON file.")
	}
}

func main() {
	if backup {
		backupFirestore()
	}
	if restore {
		restoreFirestore()
	}
}

func backupFirestore() {
	//
}

func restoreFirestore() {
	//
}

func validateJson() bool {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return json.Valid(data)
}
