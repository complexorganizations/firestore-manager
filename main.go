package main

import (
	"context"
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
	// Create a client and make sure the client is connected.
	createClient()
	// Determine what to do.
	if backup {
		backupFirestore()
	}
	if restore {
		restoreFirestore()
	}
}

func createClient() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "projectID")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	return json.Valid(data)
}
