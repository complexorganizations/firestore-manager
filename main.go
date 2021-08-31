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
	// You need to choose what to do either backup or restore.
	if !backup && !restore {
		log.Fatal("No action specified.")
	}
	// You cannot do a backup and a restore at the same time.
	if backup && restore {
		log.Fatal("Cannot backup and restore at the same time.")
	}
	// You need to specify a path to the backup file.
	if path == "example.example" && filepath.Ext(path) != ".json" {
		log.Fatal("No path specified.")
	}
	// You need to specify a valid json file.
	if !validateJson() {
		log.Fatal("Invalid JSON file.")
	}
}

func main() {
	// Determine what to do.
	if backup {
		backupFirestore()
	}
	if restore {
		restoreFirestore()
	}
}
func backupFirestore() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "projectID")
	if err != nil {
		log.Println(err)
	}
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

func openAndRead() []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
